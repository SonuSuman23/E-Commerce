package models

type Cart struct {
    ID     uint   `gorm:"primaryKey"`
    UserID uint   `gorm:"unique"`
    Items  []Item `gorm:"many2many:cart_items"`
}