package service

import (
	"github.com/google/uuid"
	dto2 "github.com/ramdanariadi/grocery-product-service/main/dto"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
	"strings"
)

type CartServiceImpl struct {
	DB *gorm.DB
}

func NewCartService(DB *gorm.DB) CartService {
	return &CartServiceImpl{DB: DB}
}

func (service CartServiceImpl) Store(productId string, total uint, userId string) *dto2.CartTotalItemDTO {
	var productRef model.Product
	tx := service.DB.Where("id = ?", productId).Find(&productRef)
	if tx.Error != nil {
		panic(exception.ValidationException{Message: "INVALID_PRODUCT"})
	}

	id, _ := uuid.NewUUID()
	saveCart := model.Cart{
		ID:        id.String(),
		UserId:    userId,
		ProductId: productRef.ID,
		Total:     total,
	}
	save := service.DB.Create(&saveCart)
	utils.PanicIfError(save.Error)

	var count int64
	tx = service.DB.Model(model.Cart{}).Where("user_id = ?", userId).Count(&count)
	utils.PanicIfError(tx.Error)

	return &dto2.CartTotalItemDTO{TotalItem: int64(uint(count))}
}

func (service CartServiceImpl) Destroy(id string, userId string) {
	cartRef := model.Cart{ID: id}
	tx := service.DB.Find(&cartRef)
	if tx.Error != nil {
		panic(exception.ValidationException{Message: "INVALID_CART"})
	}
	db := service.DB.Delete(&cartRef)
	utils.PanicIfError(db.Error)
}

func (service CartServiceImpl) Find(reqBody *dto2.FindCartDTO) []*dto2.Cart {
	var carts []*model.Cart
	tx := service.DB.Model(&carts)
	tx.Joins("LEFT JOIN products p ON p.id = carts.product_id AND p.deleted_at IS NULL")
	tx.Joins("LEFT JOIN shops s ON p.shop_id = s.id")
	tx.Joins("LEFT JOIN categories c ON p.category_id = c.id")
	tx.Preload("Product.Category").Preload("Product.Shop")
	if reqBody.Search != nil {
		tx.Where("LOWER(p.name) LIKE ?", strings.ToLower("%"+*reqBody.Search+"%"))
	}
	tx.Limit(reqBody.PageSize).Offset(reqBody.PageIndex * reqBody.PageSize).Find(&carts)
	utils.PanicIfError(tx.Error)

	result := make([]*dto2.Cart, 0)
	for _, data := range carts {
		result = append(result, &dto2.Cart{
			ID:          data.ID,
			Total:       data.Total,
			ProductId:   data.Product.ID,
			ShopId:      data.Product.Shop.ID,
			ShopName:    data.Product.Shop.Name,
			Name:        data.Product.Name,
			Description: data.Product.Description,
			ImageUrl:    data.Product.ImageUrl,
			Price:       data.Product.Price,
			PerUnit:     data.Product.PerUnit,
			Weight:      data.Product.Weight,
			Category:    data.Product.Category.Category,
		})
	}

	return result
}
