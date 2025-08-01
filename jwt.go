package utils

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

func GenerateToken(userID uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    })

    return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (uint, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil || !token.Valid {
        return 0, errors.New("invalid token")
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        userID := uint(claims["user_id"].(float64))
        return userID, nil
    }

    return 0, errors.New("invalid token claims")
}