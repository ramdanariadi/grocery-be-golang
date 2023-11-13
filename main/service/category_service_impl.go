package service

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	dto2 "github.com/ramdanariadi/grocery-product-service/main/dto"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
	"log"
)

type CategoryServiceImpl struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewCategoryServiceImpl(DB *gorm.DB, redisClient *redis.Client) CategoryService {
	return &CategoryServiceImpl{DB: DB, RedisClient: redisClient}
}

func (service CategoryServiceImpl) FindAll(pageIndex int, pageSize int) *dto2.AllCategories {
	var categories []*model.Category
	service.DB.Limit(pageSize).Offset(pageSize * pageIndex).Find(&categories)
	var count int64
	service.DB.Model(&model.Category{}).Where("deleted_at IS NULL").Count(&count)

	result := dto2.AllCategories{}
	result.Data = make([]*dto2.CategoryDTO, 0)

	for _, category := range categories {
		result.Data = append(result.Data, &dto2.CategoryDTO{Id: category.ID, Category: category.Category, ImageUrl: category.ImageUrl})
	}

	result.RecordsFiltered = len(result.Data)
	result.RecordsTotal = count
	return &result
}

func (service CategoryServiceImpl) FindById(id string) *dto2.CategoryDTO {
	var category model.Category
	var result dto2.CategoryDTO
	tx := service.DB.Where("id = ? AND deleted_at IS NULL", id).Find(&category)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	result.Id = category.ID
	result.Category = category.Category
	result.ImageUrl = category.ImageUrl
	return &result
}

func (service CategoryServiceImpl) Save(body *dto2.AddCategoryDTO) {
	id, err := uuid.NewUUID()
	utils.LogIfError(err)
	service.DB.Create(&model.Category{ID: id.String(), Category: body.Category, ImageUrl: body.ImageUrl})
}

func (service CategoryServiceImpl) Update(id string, body *dto2.AddCategoryDTO) {
	var category model.Category
	tx := service.DB.Where("id = ?", id).Find(&category)
	if tx.Error != nil {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	category.Category = body.Category
	category.ImageUrl = body.ImageUrl
	log.Printf("Id %s", id)
	save := service.DB.Save(&category)
	if save.Error != nil {
		panic(exception.InternalServerError)
	}
}

func (service CategoryServiceImpl) Delete(id string) {
	var category model.Category
	tx := service.DB.Where("id = ?", id).Find(&category)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	service.DB.Delete(&category)
}
