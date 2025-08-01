package routes

import (
    "ecommerce-backend/controllers"
    "ecommerce-backend/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/users", controllers.RegisterUser)
    r.POST("/users/login", controllers.LoginUser)
    r.GET("/users", controllers.ListUsers)

    r.POST("/items", controllers.CreateItem)
    r.GET("/items", controllers.ListItems)

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())

    auth.POST("/carts", controllers.AddToCart)
    auth.GET("/carts", controllers.GetCarts)

    auth.POST("/orders", controllers.CreateOrder)
    auth.GET("/orders", controllers.GetOrders)
}