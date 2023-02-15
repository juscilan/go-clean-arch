package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	CategoryRoutes(router)

	router.Run(":8080")
}
