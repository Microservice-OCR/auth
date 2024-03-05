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
    "go.mongodb.org/mongo-driver/bson"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {


    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        http.Error(w, "JWT secret not found", http.StatusInternalServerError)
        return
    }

    db.Connect()

    var input models.User

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Erreur lors de la lecture du corps de la requête", http.StatusInternalServerError)
        return
    }

    err = json.Unmarshal(body, &input)
    if err != nil {
        http.Error(w, "Erreur lors de la conversion du JSON", http.StatusInternalServerError)
        return
    }

    var existingUser models.User
    err = db.UserCollection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&existingUser)
    if err == nil {
        http.Error(w, "Email is already in use", http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Password hashing failed", http.StatusInternalServerError)
        return
    }

    input.Password = string(hashedPassword)

    _, err = db.UserCollection.InsertOne(context.TODO(), input)
    if err != nil {
        http.Error(w, "User creation failed", http.StatusInternalServerError)
        return
    }

	if err != nil {
        http.Error(w, "Erreur lors de la création de l'utilisateur: "+err.Error(), http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusOK)

}
