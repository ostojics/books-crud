package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ostojics/books-crud/database"
	"github.com/ostojics/books-crud/models"
	"github.com/ostojics/books-crud/routes"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	err := models.MigrateBooks(db)

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	err = models.MigrateUser(db)

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}
}

func main() {
	err := godotenv.Load(".env")
	
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	migrate(database.DB)

	router := gin.Default()
	routes.PublicRoutes(router)

	router.Run(":" + port)
}