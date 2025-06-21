package mango

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file %s", err.Error())
	}

	MongoDBUrl := os.Getenv("MONGO_URL")

	// mongo.NewClient()
	ops := options.Client().ApplyURI(MongoDBUrl)
	mc, err := mongo.Connect(ops)

	if err != nil {
		log.Fatalf("Unable to connect to Mongo database %s", err.Error())
	}

	log.Println("Connected to mongodb")

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // This is for older version of mongo client package

	// defer cancel()

	// mc.

	return mc
}

var MongoCon *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("go-auth").Collection(collectionName)

	return collection
}
