package service

import (
	dto2 "github.com/ramdanariadi/grocery-product-service/main/dto"
)

type CategoryService interface {
	FindAll(pageIndex int, pageSize int) *dto2.AllCategories
	FindById(id string) *dto2.CategoryDTO
	Save(body *dto2.AddCategoryDTO)
	Update(id string, body *dto2.AddCategoryDTO)
	Delete(id string)
}
