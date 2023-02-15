package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juscilan/go-clean-arch/internal/repositories"
	usecases "github.com/juscilan/go-clean-arch/internal/use-cases"
)

func ListCategoryController(ctx *gin.Context, repository repositories.CategoryRepositoryInterface) {

	useCase := usecases.NewListCategoryUseCase(repository)

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
