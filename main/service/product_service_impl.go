package service

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	dto2 "github.com/ramdanariadi/grocery-product-service/main/dto"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

type ProductServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewProductServiceImpl(DB *gorm.DB, redis *redis.Client) ProductService {
	return &ProductServiceImpl{DB: DB, Redis: redis}
}

func (service ProductServiceImpl) Save(userId string, requestBody *dto2.AddProductDTO) {
	var category model.Category
	tx := service.DB.Find(&category, "id = ?", requestBody.CategoryId)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	userShop := model.Shop{UserId: userId}
	find := service.DB.Find(&userShop)
	if find.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	product := model.Product{}
	id, err := uuid.NewUUID()
	utils.LogIfError(err)
	product.Shop = userShop
	product.ID = id.String()
	product.Name = requestBody.Name
	product.ImageUrl = requestBody.ImageUrl
	product.Category = category
	product.Weight = requestBody.Weight
	product.Price = requestBody.Price
	product.PerUnit = requestBody.PerUnit
	product.IsRecommended = requestBody.IsRecommended
	product.IsTop = requestBody.IsTop
	product.Description = requestBody.Description
	save := service.DB.Create(&product)
	utils.LogIfError(save.Error)
}

func (service ProductServiceImpl) FindAll(param *dto2.FindProductRequest) *dto2.FindProductResponse {
	var products []model.Product
	tx := service.DB.Model(&model.Product{})

	if param.Search != nil {
		tx.Where("LOWER(name) like ?", strings.ToLower("%"+*param.Search+"%"))
	}

	if param.IsTop != nil {
		tx.Where("is_top = ?", *param.IsTop)
	}

	if param.IsRecommendation != nil {
		tx.Where("is_recommended = ?", *param.IsRecommendation)
	}

	if param.CategoryId != nil {
		tx.Where("category_id = ?", *param.CategoryId)
	}

	var result dto2.FindProductResponse
	result.Data = make([]*dto2.ProductDTO, 0)

	var count int64
	tx.Count(&count)
	result.RecordsTotal = count
	tx.Limit(param.PageSize).Offset(param.PageIndex * param.PageSize).Preload("Category").Preload("Shop").Find(&products)

	for _, p := range products {
		result.Data = append(result.Data, &dto2.ProductDTO{
			ID:          p.ID,
			Name:        p.Name,
			ImageUrl:    p.ImageUrl,
			Category:    p.Category.Category,
			ShopId:      p.Shop.ID,
			ShopName:    p.Shop.Name,
			Weight:      p.Weight,
			Price:       p.Price,
			PerUnit:     p.PerUnit,
			Description: p.Description,
		})
	}
	result.RecordsFiltered = len(result.Data)
	return &result
}

func (service ProductServiceImpl) FindById(id string) *dto2.ProductDTO {
	var result dto2.ProductDTO
	var product model.Product
	ctx := context.Background()
	cache, err := service.Redis.Get(ctx, id).Result()
	utils.LogIfError(err)

	if cache != "" {
		err = json.Unmarshal([]byte(cache), &product)
		utils.LogIfError(err)
		log.Print("product with id " + id + " found in cache")
	} else {
		tx := service.DB.Model(&model.Product{}).Where("id = ?", id).Preload("Category").Preload("Shop").Find(&product)
		if tx.RowsAffected < 1 {
			panic(exception.ValidationException{Message: exception.BadRequest})
		}

		productByte, err := json.Marshal(product)
		utils.LogIfError(err)
		err = service.Redis.Set(ctx, product.ID, productByte, 1*time.Hour).Err()
		utils.LogIfError(err)
	}
	result.ID = product.ID
	result.ShopId = product.Shop.ID
	result.ShopName = product.Shop.Name
	result.Name = product.Name
	result.ImageUrl = product.ImageUrl
	result.Weight = product.Weight
	result.Price = product.Price
	result.PerUnit = product.PerUnit
	result.Description = product.Description
	result.Category = product.Category.Category
	return &result
}

func (service ProductServiceImpl) Update(id string, requestBody *dto2.AddProductDTO) {
	var category model.Category
	txCategory := service.DB.Find(&category, "id = ?", requestBody.CategoryId)
	if txCategory.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	var product model.Product
	tx := service.DB.Where("id = ?", id).Find(&product)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	product.Name = requestBody.Name
	product.ImageUrl = requestBody.ImageUrl
	product.Category = category
	product.Weight = requestBody.Weight
	product.Price = requestBody.Price
	product.PerUnit = requestBody.PerUnit
	product.IsRecommended = requestBody.IsRecommended
	product.IsTop = requestBody.IsTop
	product.Description = requestBody.Description
	service.DB.Save(&product)
}

func (service ProductServiceImpl) Delete(id string) {
	var product model.Product
	tx := service.DB.Where("id = ?", id).Find(&product)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	service.DB.Where("id = ?", id).Delete(&product)
}

func (service ProductServiceImpl) SetTop(id string) {
	var product model.Product
	tx := service.DB.Where("id = ?", id).Find(&product)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	product.IsTop = true
	service.DB.Save(&product)
}

func (service ProductServiceImpl) SetRecommendation(id string) {
	var product model.Product
	tx := service.DB.Where("id = ?", id).Find(&product)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	product.IsRecommended = true
	service.DB.Save(&product)
}
