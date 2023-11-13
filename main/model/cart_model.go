package model

import (
	"github.com/ramdanariadi/grocery-product-service/main/user"
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	ID        string `gorm:"primaryKey"`
	Total     uint
	ProductId string `gorm:"index"`
	Product   Product
	UserId    string `gorm:"index"`
	User      user.User
	CreatedAt time.Time      `json:"_"`
	UpdatedAt time.Time      `json:"_"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}
