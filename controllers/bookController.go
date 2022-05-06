package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ostojics/books-crud/database"
	"github.com/ostojics/books-crud/models"
)

func CreateBook(context *gin.Context) {
	var book models.Book

	err := context.BindJSON(&book)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.Create(&book).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Book successfully created", "data": book})
}

func GetBooks(context *gin.Context) {
	books := &[]models.Book{}

	err := database.DB.Find(books).Error
	
	if err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get books %s", err.Error())}) 
	return
	}

	context.JSON(http.StatusOK, gin.H{"data": books})
}
