package main

import (
    "github.com/gin-gonic/gin"
    "auth/handlers"
    "auth/db"    
	"log"
    "os"
)

func main() {
	jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Fatal("La variable d'environnement JWT_SECRET n'est pas définie")
    }
	
    db.Connect()

    r := gin.Default()

    r.POST("/login", handlers.Login)
    // Utiliser pour mes tests en local mais inutile pour le projet
    r.POST("/signup", handlers.SignUp)
    r.GET("/users", handlers.GetAllUsers)
    r.DELETE("/users", handlers.DeleteAllUsers)

    r.Run(":8080") 
}
