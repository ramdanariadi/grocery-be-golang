package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ramdanariadi/grocery-product-service/main/dto"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/service"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
	"strconv"
)

type CartControllerImpl struct {
	Service service.Service
}

func NewController(db *gorm.DB) CartController {
	return &CartControllerImpl{Service: service.NewService(db)}
}

func (controller CartControllerImpl) Store(ctx *gin.Context) {
	productId := ctx.Param("productId")
	total := ctx.Param("total")
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: "FORBIDDEN"})
	}
	totalParse, err := strconv.ParseUint(total, 0, 0)
	utils.PanicIfError(err)
	totalItemDTO := controller.Service.Store(productId, uint(totalParse), userId.(string))
	ctx.JSON(200, gin.H{"data": totalItemDTO})
}

func (controller CartControllerImpl) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: "FORBIDDEN"})
	}
	controller.Service.Destroy(id, userId.(string))
	ctx.JSON(200, gin.H{})
}

func (controller CartControllerImpl) Find(ctx *gin.Context) {
	var reqBody dto.FindCartDTO
	err := ctx.ShouldBind(&reqBody)
	utils.PanicIfError(err)
	find := controller.Service.Find(&reqBody)
	ctx.JSON(200, gin.H{"data": find})
}
