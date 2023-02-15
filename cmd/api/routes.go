package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juscilan/go-clean-arch/cmd/api/controllers"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")
	{
		categoryRoutes.POST("/", controllers.CreateCategoryController)
	}
}
