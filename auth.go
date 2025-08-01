package middleware

import (
    "ecommerce-backend/utils"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        userID, err := utils.ValidateToken(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        c.Set("userID", userID)
        c.Next()
    }
}