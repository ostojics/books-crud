package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ostojics/books-crud/database"
	"github.com/ostojics/books-crud/helpers"
	"github.com/ostojics/books-crud/models"
)

func Register(context *gin.Context) {
	var user models.User

	 err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	hashedPassword := helpers.HashPassword(user.Password)

	user.Password = hashedPassword

	err = database.DB.Create(&user).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}