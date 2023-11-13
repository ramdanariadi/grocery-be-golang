package controller

import (
	"github.com/gin-gonic/gin"
	dto2 "github.com/ramdanariadi/grocery-product-service/main/dto"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/service"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
)

type TransactionControllerImpl struct {
	Service service.TransactionService
}

func NewTransactionController(DB *gorm.DB) TransactionController {
	return TransactionControllerImpl{Service: &service.TransactionServiceImpl{DB: DB}}
}

func (controller TransactionControllerImpl) Save(ctx *gin.Context) {
	var request dto2.AddTransactionDTO
	ctx.ShouldBind(&request)
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: "UNAUTHORIZED"})
	}
	controller.Service.Save(&request, userId.(string))
	ctx.JSON(200, gin.H{})
}

func (controller TransactionControllerImpl) Find(ctx *gin.Context) {
	var request dto2.FindTransactionDTO
	err := ctx.ShouldBind(&request)
	utils.PanicIfError(err)
	transactionDTO := controller.Service.Find(&request)
	ctx.JSON(200, gin.H{"data": transactionDTO})
}
