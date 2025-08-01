package config

import (
    "ecommerce-backend/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    db, err := gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect database")
    }

    db.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.Order{})
    DB = db
}