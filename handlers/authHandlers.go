package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "time"
    "github.com/dgrijalva/jwt-go"
    "auth/models"
    "auth/db"
	"context"
	"os"
)

func Login(c *gin.Context) {
    var user models.User
    var foundUser models.User

    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := db.Collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&foundUser)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Identifiants incorrects"})
        return
    }

	secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "La clé secrète JWT n'est pas définie"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": user.Email,
        "exp":   time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString([]byte(secret))

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création du token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
