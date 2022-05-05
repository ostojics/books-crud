package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ostojics/books-crud/controllers"
	"github.com/ostojics/books-crud/middleware"
)

func PrivateRoutes(router *gin.Engine) {
	router.Use(middleware.Authenticate)
	router.GET("/books", controllers.GetBooks)
}