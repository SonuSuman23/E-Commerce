package controllers

import (
    "ecommerce-backend/config"
    "ecommerce-backend/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateOrder(c *gin.Context) {
    userID := c.GetUint("userID")

    var cart models.Cart
    result := config.DB.Preload("Items").Where("user_id = ?", userID).First(&cart)
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
        return
    }

    order := models.Order{UserID: userID, Items: cart.Items}
    config.DB.Create(&order)

    config.DB.Model(&cart).Association("Items").Clear()
    config.DB.Delete(&cart)

    c.JSON(http.StatusCreated, gin.H{"message": "Order created"})
}

func GetOrders(c *gin.Context) {
    userID := c.GetUint("userID")

    var orders []models.Order
    config.DB.Preload("Items").Where("user_id = ?", userID).Find(&orders)
    c.JSON(http.StatusOK, orders)
}