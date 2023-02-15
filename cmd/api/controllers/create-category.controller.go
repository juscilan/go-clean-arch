package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/juscilan/go-clean-arch/internal/use-cases"
)

type createCategotyInputDTO struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategoryController(ctx *gin.Context) {
	var body createCategotyInputDTO

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"sucess": false,
			"error":  err.Error(),
		})
		return
	}

	useCase := usecases.NewCreateCategoryUseCase()
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
