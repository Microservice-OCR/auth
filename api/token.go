package api

import (
    "context"
    "log"
	"fmt"
    "auth/db"
	"encoding/json"
	"net/http"
    "auth/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllToken(w http.ResponseWriter, r *http.Request) {
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

	jsonData, err := json.Marshal(tokens)
    if err != nil {
        http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
        return
    }

    w.Write(jsonData)
}