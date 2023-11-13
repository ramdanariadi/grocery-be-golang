package service

import (
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/dto"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
	"strings"
)

type WishlistServiceImpl struct {
	DB *gorm.DB
}

func NewWishlistServiceImpl(DB *gorm.DB) WishlistService {
	return &WishlistServiceImpl{DB: DB}
}

func (service WishlistServiceImpl) Store(productId string, userId string) {
	var p model.Product
	tx := service.DB.Where("id = ?", productId).Find(&p)
	if tx.Error != nil {
		panic(exception.ValidationException{Message: "INVALID_PRODUCT"})
	}
	id, _ := uuid.NewUUID()
	wishlist := model.Wishlist{
		ID:      id.String(),
		Product: p,
		UserId:  userId,
	}
	save := service.DB.Create(&wishlist)
	utils.PanicIfError(save.Error)
}

func (service WishlistServiceImpl) Destroy(productId string, userId string) {
	wishlist := model.Wishlist{ProductId: productId, UserId: userId}
	find := service.DB.Find(&wishlist)
	if find.Error != nil {
		panic(exception.ValidationException{"INVALID_WISHLIST"})
	}
	tx := service.DB.Delete(&wishlist)
	utils.PanicIfError(tx.Error)
}

func (service WishlistServiceImpl) Find(reqBody *dto.FindWishlistDTO) []*dto.WishlistDTO {
	var wishlists []*model.Wishlist
	tx := service.DB.Model(&model.Wishlist{})
	tx.Joins("LEFT JOIN products p ON p.id = wishlists.product_id AND p.deleted_at IS NULL")
	tx.Joins("LEFT JOIN categories c ON c.id = p.category_id AND c.deleted_at IS NULL")
	tx.Preload("Product.Category")
	if reqBody.Search != nil {
		tx.Where("LOWER(p.name) like ?", strings.ToLower("%"+*reqBody.Search+"%"))
	}
	tx.Limit(reqBody.PageSize).Offset(reqBody.PageIndex * reqBody.PageSize).Find(&wishlists)
	wishlistsResult := make([]*dto.WishlistDTO, 0)
	for _, wishlist := range wishlists {
		wishlistsResult = append(wishlistsResult, &dto.WishlistDTO{
			ID:          wishlist.ID,
			ProductId:   wishlist.ProductId,
			Name:        wishlist.Product.Name,
			Category:    wishlist.Product.Category.Category,
			ImageUrl:    wishlist.Product.ImageUrl,
			Price:       wishlist.Product.Price,
			PerUnit:     wishlist.Product.PerUnit,
			Weight:      wishlist.Product.Weight,
			Description: wishlist.Product.Description,
		})
	}
	return wishlistsResult
}

func (service WishlistServiceImpl) FindByProductId(productId string, userId string) *dto.WishlistDTO {
	wishlist := model.Wishlist{ProductId: productId, UserId: userId}
	find := service.DB.Model(&model.Wishlist{}).Where("product_id = ? AND user_id = ?", productId, userId).Preload("Product.Category").Find(&wishlist)
	if find.RowsAffected < 1 {
		panic(exception.ValidationException{"INVALID_WISHLIST"})
	}

	wishlistDTO := &dto.WishlistDTO{
		ID:          wishlist.ID,
		ProductId:   wishlist.ProductId,
		Name:        wishlist.Product.Name,
		Category:    wishlist.Product.Category.Category,
		ImageUrl:    wishlist.Product.ImageUrl,
		Price:       wishlist.Product.Price,
		PerUnit:     wishlist.Product.PerUnit,
		Weight:      wishlist.Product.Weight,
		Description: wishlist.Product.Description,
	}

	return wishlistDTO
}
