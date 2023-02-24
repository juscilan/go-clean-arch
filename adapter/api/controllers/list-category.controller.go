package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	aplication "github.com/juscilan/go-clean-arch/internal/application"
	"github.com/juscilan/go-clean-arch/internal/repositories/interfaces"
)

func ListCategoryController(ctx *gin.Context, repository interfaces.CategoryRepositoryInterface) {

	useCase := aplication.NewListCategoryUseCase(repository)

	categories, err := useCase.Execute()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"sucess": false,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"sucess": true, "categories": categories})
}
