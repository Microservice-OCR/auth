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
    "log"
    "golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
    var user models.User
    var foundUser models.User

    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := db.UserCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&foundUser)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Identifiants incorrects"})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
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
        "exp":   time.Now().Add(time.Minute * 30).Unix(),
    })

    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création du token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func SignUp(c *gin.Context) {
    var user models.User

    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de hasher le mot de passe"})
        return
    }
    user.Password = string(hashedPassword)

    _, err = db.UserCollection.InsertOne(context.TODO(), user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de l'utilisateur"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Utilisateur créé avec succès"})
}

func GetAllUsers(c *gin.Context) {
    var users []models.User

    cursor, err := db.UserCollection.Find(context.TODO(), bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des utilisateurs"})
        return
    }

    for cursor.Next(context.TODO()) {
        var user models.User
        err := cursor.Decode(&user)
        if err != nil {
            log.Fatal(err)
        }
        users = append(users, user)
    }

    if err := cursor.Err(); err != nil {
        log.Fatal(err)
    }

    cursor.Close(context.TODO())

    c.JSON(http.StatusOK, users)
}

func DeleteAllUsers(c *gin.Context) {
    // Supprimer tous les documents de la collection
    _, err := db.UserCollection.DeleteMany(context.TODO(), bson.D{{}})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression des utilisateurs"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Tous les utilisateurs ont été supprimés avec succès"})
}
