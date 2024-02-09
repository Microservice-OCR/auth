package main

import (
    "github.com/gin-gonic/gin"
    "auth/handlers"
    "auth/db"    
	"log"
    "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Erreur lors du chargement du fichier .env")
    }
	
    db.Connect()

    r := gin.Default()

    r.POST("/login", handlers.Login)

    r.Run(":8080") 
}
