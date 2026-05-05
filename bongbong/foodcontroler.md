You’ve got a mix of small syntax mistakes, MongoDB driver version mismatches, and a couple of missing helper functions. Let’s go through the issues one by one and fix them cleanly.

---

## 1. ❌ `undefined: inTimeSpan`

You never defined this function. Add it:

```go
func inTimeSpan(start, end, check time.Time) bool {
	return (check.After(start) && check.Before(end)) || check.Equal(start) || check.Equal(end)
}
```

---

## 2. ❌ `bson.E` vs `primitive.E` mismatch

You’re mixing:

```go
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/v2/bson"
```

These are **different packages**, and their types are incompatible.

### ✅ Fix:

Pick ONE BSON package (recommended: v2 since you're already using it)

Replace:

```go
var updateObj primitive.D
```

with:

```go
var updateObj bson.D
```

---

## 3. ❌ Unkeyed struct fields in `bson.E`

Wrong:

```go
bson.E{"start_date", menu.Start_Date}
```

Correct:

```go
bson.E{Key: "start_date", Value: menu.Start_Date}
```

Do this for all `bson.E`.

---

## 4. ❌ Wrong field name in filter

You wrote:

```go
filter := bson.M{"menuu_id": menuId}
```

Fix typo:

```go
filter := bson.M{"menu_id": menuId}
```

---

## 5. ❌ `options.UpdateOptions` field name

Wrong:

```go
opt := options.UpdateOptions{
	upsert: &upsert,
}
```

Correct:

```go
opt := options.UpdateOptions{
	Upsert: &upsert,
}
```

---

## 6. ❌ Invalid UpdateOne syntax

This is completely broken:

```go
menu.Collection.UpdateOne{
	ctx,
	filter,
	bson.D{
		{"$sets", updateObj},
	},
	&opt,
}
```

### ✅ Fix:

```go
result, err := menuCollection.UpdateOne(
	ctx,
	filter,
	bson.D{{Key: "$set", Value: updateObj}},
	&opt,
)
```

---

## 7. ❌ `undefined: err` and `undefined: result`

You used `err` and `result` without declaring them.

Fix already included above:

```go
result, err := menuCollection.UpdateOne(...)
```

---

## 8. ❌ Typo in `$set`

You wrote:

```go
"$sets"
```

Correct:

```go
"$set"
```

---

## 9. ❌ Broken field name: `Catagory`

You used:

```go
menu.Catagory
```

Then:

```go
menu.Category
```

Pick ONE (probably `Category`).

---

## 10. ❌ Missing `round` and `toFixed`

Add these:

```go
import "math"

func round(num float64) int {
	return int(math.Round(num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
```

---

## 11. ⚠️ Minor logic issue

This line:

```go
*&menu.Start_Date
```

Is unnecessary and confusing.

Fix:

```go
inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now())
```

---

## ✅ Cleaned `UpdateFood` (fixed version)

```go
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		var updateObj bson.D

		if menu.Start_Date != nil && menu.End_Date != nil {
			if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid time range"})
				return
			}

			updateObj = append(updateObj, bson.E{Key: "start_date", Value: menu.Start_Date})
			updateObj = append(updateObj, bson.E{Key: "end_date", Value: menu.End_Date})
		}

		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
		}

		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{Key: "category", Value: menu.Category})
		}

		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: menu.Updated_at})

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := menuCollection.UpdateOne(
			ctx,
			filter,
			bson.D{{Key: "$set", Value: updateObj}},
			&opt,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu update failed"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
```

---

## Big Picture Problem ⚠️

You're mixing **MongoDB driver v1 and v2 APIs**, which causes half your errors.

### Strong recommendation:

Stick to ONE:

* Either: `go.mongodb.org/mongo-driver` (stable, most tutorials)
* Or: `go.mongodb.org/mongo-driver/v2` (newer, but less common)

Right now you're mixing both — that’s the root cause of many issues.

---

If you want, I can clean your entire controller file so it's production-ready and consistent.
