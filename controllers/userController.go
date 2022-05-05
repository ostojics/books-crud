package controllers

import (
	"fmt"
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
		fmt.Println(err)
		return
	}
	
	hashedPassword := helpers.HashPassword(user.Password)

	user.Password = hashedPassword

	err = database.DB.Create(&user).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func Login(context *gin.Context) {
	var user models.User
	foundUser := &models.User{}

	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.Where("email = ?", user.Email).First(foundUser).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not get user with this email"})
		fmt.Println(err)
		return
	}

	valid, msg := helpers.VerifyPassword(foundUser.Password, user.Password)

	if !valid {
		context.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	token, err := helpers.CreateToken(user.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.Writer.Header().Set("Authorization", token)
	context.JSON(http.StatusOK, gin.H{"data": foundUser})
}