package controllers

import (
    "ecommerce-backend/config"
    "ecommerce-backend/models"
    "ecommerce-backend/utils"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "net/http"
)

func RegisterUser(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    input.Password = string(hashed)

    config.DB.Create(&input)
    c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func LoginUser(c *gin.Context) {
    var input models.User
    var dbUser models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Where("username = ?", input.Username).First(&dbUser)
    if dbUser.ID == 0 || bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password)) != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, _ := utils.GenerateToken(dbUser.ID)
    c.JSON(http.StatusOK, gin.H{"token": token})
}

func ListUsers(c *gin.Context) {
    var users []models.User
    config.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}