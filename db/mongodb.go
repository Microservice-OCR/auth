package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "os"
)

var UserCollection *mongo.Collection
var TokenCollection *mongo.Collection

func Connect() {
    uri := os.Getenv("MONGO_URI")
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    if err != nil {
        log.Println("Erreur lors de la connexion à la base de données")
        log.Fatal(err)
    }

    UserCollection = client.Database("auth").Collection("users")
    TokenCollection = client.Database("auth").Collection("tokens")
}
