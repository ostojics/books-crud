package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/ostojics/books-crud/controllers"
)

func PublicRoutes(router *gin.Engine) {
	router.POST("/users/register", controllers.Register)
	router.POST("users/login", controllers.Login)
	router.GET("/books", controllers.GetBooks)
}