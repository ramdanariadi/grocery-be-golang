package service

import (
	"encoding/json"
	"github.com/google/uuid"
	dto2 "github.com/ramdanariadi/grocery-product-service/main/dto"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	model2 "github.com/ramdanariadi/grocery-product-service/main/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
	"log"
)

type TransactionServiceImpl struct {
	DB *gorm.DB
}

type Collection []string

func (collection Collection) isExist(item string) bool {
	for _, s := range collection {
		if s == item {
			return true
		}
	}
	return false
}

func (service TransactionServiceImpl) Save(request *dto2.AddTransactionDTO, userId string) {
	marshal, _ := json.Marshal(request)
	log.Println("request body " + string(marshal))
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var productIds Collection
		for _, item := range request.Data {
			if !productIds.isExist(item.ProductId) {
				productIds = append(productIds, item.ProductId)
			}
		}

		var products []*model2.Product
		tx.Model(&model2.Product{}).Where("id IN ?", productIds).Preload("Category").Find(&products)

		if len(products) != len(productIds) {
			panic(exception.ValidationException{Message: "INVALID_PRODUCT"})
		}

		productMap := map[string]*model2.Product{}
		var totalPrice uint64
		for _, p := range products {
			totalPrice += p.Price
			productMap[p.ID] = p
		}

		id, _ := uuid.NewUUID()
		transaction := model2.Transaction{
			ID:         id.String(),
			UserId:     userId,
			TotalPrice: totalPrice,
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}
		var transactionDetails []*model2.TransactionDetail
		for _, d := range request.Data {
			p := productMap[d.ProductId]
			dtId, _ := uuid.NewUUID()
			detail := model2.TransactionDetail{ID: dtId.String(), Transaction: transaction, Product: *p, Total: d.Total, Name: p.Name, Price: p.Price, ImageUrl: p.ImageUrl, Description: p.Description, PerUnit: p.PerUnit, Weight: p.Weight, CategoryId: p.CategoryId, Category: p.Category.Category}
			transactionDetails = append(transactionDetails, &detail)
		}
		log.Printf("request body data length %d", len(transactionDetails))
		if err := tx.Create(&transactionDetails).Error; err != nil {
			return err
		}

		if db := tx.Where("user_id = ?", userId).Delete(&model2.Cart{}); nil != db.Error {
			return db.Error
		}

		log.Println("success Save detail transaction")
		return nil
	})
	utils.PanicIfError(err)
}

func (service TransactionServiceImpl) Find(param *dto2.FindTransactionDTO) []*dto2.TransactionDTO {
	var transactions []*model2.Transaction
	tx := service.DB.Model(&model2.Transaction{})
	tx.Joins("LEFT JOIN transaction_details td ON td.transaction_id = transactions.id")
	tx.Joins("LEFT JOIN products p ON td.product_id = p.id AND p.deleted_at IS NULL")
	tx.Preload("TransactionDetails.Product")
	if param.Search != nil {
		tx.Where("LOWER(p.name) ilike ?", "%"+*param.Search+"%")
	}
	tx.Limit(param.PageSize).Offset(param.PageIndex * param.PageSize).Find(&transactions)

	result := make([]*dto2.TransactionDTO, 0)
	for _, t := range transactions {
		transactionDTO := dto2.TransactionDTO{Id: t.ID, PriceTotal: 0}
		for _, td := range t.TransactionDetails {
			p := td.Product
			item := dto2.TransactionItemDTO{ID: td.ID, Name: td.Name, Price: td.Price, PerUnit: td.PerUnit, Weight: td.Weight, ImageUrl: p.ImageUrl, Description: p.Description, Total: td.Total}
			transactionDTO.Items = append(transactionDTO.Items, &item)
			transactionDTO.PriceTotal += td.Price
		}
		result = append(result, &transactionDTO)
	}

	return result
}
