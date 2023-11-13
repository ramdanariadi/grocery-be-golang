package service

import (
	dto2 "github.com/ramdanariadi/grocery-product-service/main/dto"
)

type ProductService interface {
	Save(userId string, product *dto2.AddProductDTO)
	FindAll(param *dto2.FindProductRequest) *dto2.FindProductResponse
	FindById(id string) *dto2.ProductDTO
	Update(id string, product *dto2.AddProductDTO)
	Delete(id string)
	SetTop(id string)
	SetRecommendation(id string)
}
