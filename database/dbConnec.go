package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	log.Println("Connected to MongoDB")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("restaurant").Collection(collectionName)
}














// package database

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/v2/mongo"
// 	"go.mongodb.org/mongo-driver/v2/mongo/options"
// )

// func DBinstance() *mongo.Client {
// 	MongoDb := "mongodb://localhost:27017"
// 	fmt.Print(MongoDb)

// 	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	err = client.Connect(ctx)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("connected to mongoBD")

// 	return client
// }

// var Client *mongo.Client = DBinstance()

// func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
// 	var collection *mongo.Collection = client.Database("resturant").Collection(collectionName)

// 	return collection
// }
