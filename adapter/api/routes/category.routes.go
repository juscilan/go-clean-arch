package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juscilan/go-clean-arch/adapter/api/controllers"
	inmemorydb "github.com/juscilan/go-clean-arch/adapter/db/in-memory-db"
)

func CategoryRoutes(router *gin.Engine) {
	inMemoryCategoryRepository := inmemorydb.NewCategoryInMemoryRepository()

	rg := router.Group("/categories")

	rg.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategoryController(ctx, inMemoryCategoryRepository)
	})

	rg.GET("/", func(ctx *gin.Context) {
		controllers.ListCategoryController(ctx, inMemoryCategoryRepository)
	})
}
