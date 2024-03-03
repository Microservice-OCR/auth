package api

import (
	"auth/db"
	"auth/models"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
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