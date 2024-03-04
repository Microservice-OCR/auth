package api

import (
    "auth/db"
    "auth/models"
    "context"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "os"
    "time"

    "github.com/dgrijalva/jwt-go"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        http.Error(w, "JWT secret not found", http.StatusInternalServerError)
        return
    }

    db.Connect()

    var foundUser models.User
    input := models.User{}

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

    err = db.Collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&foundUser)
    if err != nil {
        http.Error(w, "Incorrect credentials", http.StatusUnauthorized)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(input.Password))
    if err != nil {
        http.Error(w, "Incorrect credentials", http.StatusUnauthorized)
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": input.Email,
        "exp":   time.Now().Add(time.Minute * 30).Unix(),
    })

    tokenString, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        http.Error(w, "JWT creation failed", http.StatusInternalServerError)
        return
    }

    _, err = db.Collection.InsertOne(context.TODO(), bson.M{
        "token": tokenString,
    })
    if err != nil {
        http.Error(w, "Erreur lors de l'enregistrement du token en base de données", http.StatusInternalServerError)
        return
    }

    tokenOutput := models.Output{
        Token:       tokenString,
        ConnectedAt: time.Now(),
    }

    jsonData, err := json.Marshal(tokenOutput)
    if err != nil {
        http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
        return
    }

    w.Write(jsonData)
}