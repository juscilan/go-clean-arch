package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juscilan/go-clean-arch/cmd/api/controllers"
	"github.com/juscilan/go-clean-arch/internal/repositories"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	inMemoryCategoryRepository := repositories.NewCategoryInMemoryRepository()

	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategoryController(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.GET("/", func(ctx *gin.Context) {
		controllers.ListCategoryController(ctx, inMemoryCategoryRepository)
	})

}
