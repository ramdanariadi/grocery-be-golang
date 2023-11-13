package model

import (
	"gorm.io/gorm"
	"time"
)

type Wishlist struct {
	ID        string `gorm:"primaryKey"`
	ProductId string `gorm:"index"`
	Product   Product
	UserId    string         `gorm:"index"`
	CreatedAt time.Time      `json:"_"`
	UpdatedAt time.Time      `json:"_"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}
