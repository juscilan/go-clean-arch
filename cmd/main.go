package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juscilan/go-clean-arch/adapter/api/routes"
)

func main() {
	router := gin.Default()
	routes.CategoryRoutes(router)
	router.Run(":8080")
}
