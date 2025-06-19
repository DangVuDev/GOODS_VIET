package config

import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() {
    clientOptions := options.Client().ApplyURI("mongodb+srv://bachaidang19082004:zhiox9S6SYd6spm7@hubtsocial.7k7uj.mongodb.net/")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    Client = client
}