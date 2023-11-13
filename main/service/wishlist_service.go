package service

import (
	dto2 "github.com/ramdanariadi/grocery-product-service/main/dto"
)

type WishlistService interface {
	Store(productId string, userId string)
	Destroy(productId string, userId string)
	Find(reqBody *dto2.FindWishlistDTO) []*dto2.WishlistDTO
	FindByProductId(productId string, userId string) *dto2.WishlistDTO
}
