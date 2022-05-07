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
	var existingBook models.Book

	err := context.BindJSON(&book)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_ = database.DB.Where("title = ?", book.Title).First(&existingBook).Error

	if existingBook.Title != "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "book with this name already exists"})
		return
	}

	err = database.DB.Create(&book).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "book successfully created", "data": book})
}

func GetBooks(context *gin.Context) {
	books := &[]models.Book{}

	err := database.DB.Find(books).Error
	
	if err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get books %s", err.Error())}) 
	return
	}

	context.JSON(http.StatusOK, gin.H{"data": books})
}

func GetBookById(context *gin.Context) {
	id := context.Param("id")
	book := &models.Book{}

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id can't be empty"})
		return
	}

	err := database.DB.Where("id = ?", id).First(book).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": book})
}
