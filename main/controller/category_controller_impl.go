package controller

import (
	"github.com/gin-gonic/gin"
	dto2 "github.com/ramdanariadi/grocery-product-service/main/dto"
	"github.com/ramdanariadi/grocery-product-service/main/service"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
)

type CategoryControllerImpl struct {
	Service service.CategoryService
}

func NewCategoryController(db *gorm.DB) CategoryController {
	return &CategoryControllerImpl{
		Service: service.NewCategoryServiceImpl(db, nil),
	}
}

func (controller CategoryControllerImpl) FindAll(ctx *gin.Context) {
	var param dto2.PaginationDTO
	err := ctx.ShouldBindQuery(&param)
	utils.PanicIfError(err)
	ctx.JSON(200, controller.Service.FindAll(param.PageIndex, param.PageSize))
}

func (controller CategoryControllerImpl) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	category := controller.Service.FindById(id)
	ctx.JSON(200, gin.H{"data": category})
}

func (controller CategoryControllerImpl) Save(ctx *gin.Context) {
	request := dto2.AddCategoryDTO{}
	err := ctx.Bind(&request)
	utils.PanicIfError(err)
	controller.Service.Save(&request)
	ctx.JSON(200, gin.H{})
}

func (controller CategoryControllerImpl) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var request dto2.AddCategoryDTO
	ctx.Bind(&request)
	controller.Service.Update(id, &request)
	ctx.JSON(200, gin.H{})
}

func (controller CategoryControllerImpl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	controller.Service.Delete(id)
	ctx.JSON(200, gin.H{})
}
