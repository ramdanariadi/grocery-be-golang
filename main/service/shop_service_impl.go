package service

import (
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/dto"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/model"
	userModel "github.com/ramdanariadi/grocery-product-service/main/user"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
	"log"
)

type ShopServiceImpl struct {
	*gorm.DB
}

func (service *ShopServiceImpl) AddShop(userId string, shop dto.AddShopDTO) {
	if nil == shop.Address || nil == shop.Name {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	var user userModel.User
	find := service.DB.Where("id = ?", userId).Find(&user)
	if nil != find.Error {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	newUUID, _ := uuid.NewUUID()
	s := model.Shop{
		ID:       newUUID.String(),
		Name:     *shop.Name,
		Address:  *shop.Address,
		User:     user,
		ImageUrl: *shop.ImageUrl,
	}
	tx := service.DB.Create(&s)
	utils.PanicIfError(tx.Error)
}

func (service *ShopServiceImpl) UpdateShop(userId string, reqBody dto.EditShopDTO) {
	var shop model.Shop
	tx := service.DB.Where("user_id = ? ", userId).Find(&shop)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	shop.Name = *reqBody.Name
	shop.Address = *reqBody.Address
	shop.ImageUrl = *reqBody.ImageUrl

	save := service.DB.Save(&shop)
	utils.PanicIfError(save.Error)
}

func (service *ShopServiceImpl) GetShop(userId string) dto.ShopDTO {
	var shop model.Shop
	log.Print("user id : " + userId)
	tx := service.DB.Where("user_id = ?", userId).Find(&shop)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	return dto.ShopDTO{Id: shop.ID, Name: shop.Name, Address: shop.Address, ImageUrl: shop.ImageUrl}
}

func (service *ShopServiceImpl) DeleteShop(userID string) {
	var shop model.Shop
	tx := service.DB.Where("user_id = ?", userID).Find(&shop)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	service.DB.Delete(&shop)
}
