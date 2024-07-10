package database


import (
	"context"
	"log"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func Connect(){
	var mongoUri = os.Getenv("MONGO_URI")
	client,err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err!=nil{
		log.Fatal(err);
	}
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel();
	err = client.Connect(ctx);
	if err!=nil{
		log.Fatal(err);
	}
	DB = client;
	log.Println("Connected To MongoDB");
}

