package service

import (
	dto2 "github.com/ramdanariadi/grocery-product-service/main/dto"
)

type CartService interface {
	Store(productId string, total uint, userId string) *dto2.CartTotalItemDTO
	Destroy(id string, userId string)
	Find(reqBody *dto2.FindCartDTO) []*dto2.Cart
}
