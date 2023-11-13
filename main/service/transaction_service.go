package service

import (
	dto2 "github.com/ramdanariadi/grocery-product-service/main/dto"
)

type TransactionService interface {
	Save(request *dto2.AddTransactionDTO, userId string)
	Find(param *dto2.FindTransactionDTO) []*dto2.TransactionDTO
}
