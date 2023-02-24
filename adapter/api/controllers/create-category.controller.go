package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	aplication "github.com/juscilan/go-clean-arch/internal/application"
	"github.com/juscilan/go-clean-arch/internal/repositories/interfaces"
)

type createCategotyInputDTO struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategoryController(ctx *gin.Context, repository interfaces.CategoryRepositoryInterface) {
	var body createCategotyInputDTO

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"sucess": false,
			"error":  err.Error(),
		})
		return
	}

	useCase := aplication.NewCreateCategoryUseCase(repository)

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
