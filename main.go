package main

import (
    "ecommerce-backend/config"
    "ecommerce-backend/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    config.ConnectDatabase()
    routes.SetupRoutes(r)

    r.Run(":8080")
}