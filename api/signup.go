package api

import (
	"auth/db"
	"auth/models"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : SUPPRIMER POUR PROD
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*") // or specify your domain
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	
	jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
		http.Error(w, "JWT secret not found", http.StatusInternalServerError)
    }
	
    db.Connect()
	
	input := models.User{}
	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du corps de la requête", http.StatusInternalServerError)
		return
	}

	// Décodage du JSON dans la structure IInput
	err = json.Unmarshal(body, &input)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion du JSON", http.StatusInternalServerError)
		return
	}

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
		http.Error(w, "Password hashing failed", http.StatusInternalServerError)
        return
    }

    input.Password = string(hashedPassword)

    _, err = db.Collection.InsertOne(context.TODO(), input)
    if err != nil {
		http.Error(w, "User creation failed", http.StatusInternalServerError)
        return
    }
	w.Write([]byte("User created successfully"))
}