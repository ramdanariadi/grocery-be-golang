package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ramdanariadi/grocery-product-service/main/dto"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/service"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
)

type WishlistControllerImpl struct {
	Service service.WishlistService
}

func NewWishlistController(db *gorm.DB) WishlistController {
	return &WishlistControllerImpl{Service: service.NewWishlistServiceImpl(db)}
}

func (controller WishlistControllerImpl) Store(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: "UNAUTHORIZED"})
	}
	productId := ctx.Param("productId")
	controller.Service.Store(productId, userId.(string))
	ctx.JSON(200, gin.H{})
}

func (controller WishlistControllerImpl) Destroy(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: "UNAUTHORIZED"})
	}
	productId := ctx.Param("productId")
	controller.Service.Destroy(productId, userId.(string))
	ctx.JSON(200, gin.H{})
}

func (controller WishlistControllerImpl) Find(ctx *gin.Context) {
	var reqBody dto.FindWishlistDTO
	err := ctx.ShouldBind(&reqBody)
	utils.PanicIfError(err)
	wishlists := controller.Service.Find(&reqBody)
	ctx.JSON(200, gin.H{"data": wishlists})
}

func (controller WishlistControllerImpl) FindByProductId(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: "UNAUTHORIZED"})
	}
	productId := ctx.Param("productId")
	wishlist := controller.Service.FindByProductId(productId, userId.(string))
	ctx.JSON(200, gin.H{"data": wishlist})
}
