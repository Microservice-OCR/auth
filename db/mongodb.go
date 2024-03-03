package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "os"
)

var Collection *mongo.Collection

func Connect() {
    uri := os.Getenv("MONGO_URI")
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    if err != nil {
        log.Println("Erreur lors de la connexion à la base de données")
        log.Fatal(err)
    }

    Collection = client.Database("auth").Collection("users")
}
