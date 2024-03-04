package api

import (
    "net/http"
    "strings"
	"auth/db"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "log"
	// "encoding/json"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    tokenHeader := r.Header.Get("Authorization")
    if tokenHeader == "" {
        http.Error(w, "No token provided", http.StatusBadRequest)
        return
    }

    token := strings.TrimPrefix(tokenHeader, "Bearer ")
    if token == "" {
        http.Error(w, "No token provided", http.StatusBadRequest)
        return
    }

    db.Connect()

	_, err := db.Collection.DeleteOne(context.TODO(), bson.M{"token": token})
	if err != nil {
		http.Error(w, "Error deleting token", http.StatusInternalServerError)
		return
	}

    w.Write([]byte("User logged out successfully"))
}


