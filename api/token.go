package api

import (
    "context"
    "log"
    "auth/db"
    "auth/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllToken() ([]models.Output, error) {
	db.Connect()

	var tokens []models.Output
	cursor, err := db.Collection.Find(context.Background(), bson.D{{}}, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var token models.Output
		cursor.Decode(&token)
		tokens = append(tokens, token)
	}

	return tokens, nil
}