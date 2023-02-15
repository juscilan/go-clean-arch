package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juscilan/go-clean-arch/internal/repositories"
	usecases "github.com/juscilan/go-clean-arch/internal/use-cases"
)

type createCategotyInputDTO struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategoryController(ctx *gin.Context, repository repositories.CategoryRepositoryInterface) {
	var body createCategotyInputDTO

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"sucess": false,
			"error":  err.Error(),
		})
		return
	}

	useCase := usecases.NewCreateCategoryUseCase(repository)

	err = useCase.Execute(body.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"sucess": false,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"sucess": true})
}
