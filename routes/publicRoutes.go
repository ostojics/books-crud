package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/ostojics/books-crud/controllers"
)

func PublicRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
}