package controllers

import (
    "ecommerce-backend/config"
    "ecommerce-backend/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateItem(c *gin.Context) {
    var item models.Item
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Create(&item)
    c.JSON(http.StatusCreated, item)
}

func ListItems(c *gin.Context) {
    var items []models.Item
    config.DB.Find(&items)
    c.JSON(http.StatusOK, items)
}