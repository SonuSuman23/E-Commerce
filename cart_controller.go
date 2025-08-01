package controllers

import (
    "ecommerce-backend/config"
    "ecommerce-backend/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

type AddToCartRequest struct {
    ItemIDs []uint `json:"item_ids"`
}

func AddToCart(c *gin.Context) {
    userID := c.GetUint("userID")
    var req AddToCartRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var cart models.Cart
    config.DB.Where("user_id = ?", userID).Preload("Items").FirstOrCreate(&cart, models.Cart{UserID: userID})

    var items []models.Item
    config.DB.Where("id IN ?", req.ItemIDs).Find(&items)
    config.DB.Model(&cart).Association("Items").Append(items)

    c.JSON(http.StatusOK, gin.H{"message": "Items added to cart"})
}

func GetCarts(c *gin.Context) {
    var carts []models.Cart
    config.DB.Preload("Items").Find(&carts)
    c.JSON(http.StatusOK, carts)
}