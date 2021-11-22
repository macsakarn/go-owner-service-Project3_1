package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb)) //ขั้นตอนนี้เป็นแค่การสร้าง client ไว้เท่านั้น ไม่ได้เชื่อมต่อจริงๆ
	if err != nil {
		log.Fatal(err)
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second) //time Out
	defer cancel()
	err = client.Connect(ctx) //คำสั่งในการ Connect
	if err != nil {
		log.Fatal(err)
	}

	println("Database Connect!!!!!")

	return client
}

var Client *mongo.Client = ConnectDB()

func GetMongoDbCollection(client *mongo.Client, CollectionName string) *mongo.Collection {

	collection := client.Database("mobileDB").Collection(CollectionName)

	return collection
}

func Disconnect(client *mongo.Client) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second) //time Out
	defer cancel()
	client.Disconnect(ctx)
	println("Database Disconnect")
}
