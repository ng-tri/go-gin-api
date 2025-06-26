package main

import (
    "github.com/gin-gonic/gin"
    "go_gin_api/routes"
)

func main() {
    r := gin.Default()

    routes.RegisterProductRoutes(r)

    r.Run(":8080")
}
