package main

import (
	"go_gin_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterProductRoutes(r)

	r.Run(":8080")
}
