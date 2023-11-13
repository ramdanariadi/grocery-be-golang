package service

import "github.com/ramdanariadi/grocery-product-service/main/dto"

type ShopService interface {
	AddShop(userId string, shop dto.AddShopDTO)
	UpdateShop(userId string, shop dto.EditShopDTO)
	GetShop(userId string) dto.ShopDTO
	DeleteShop(userID string)
}
