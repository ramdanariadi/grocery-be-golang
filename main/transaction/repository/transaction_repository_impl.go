package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/transaction/model"
	"log"
	"strings"
	"time"
)

type TransactionRepositoryImpl struct {
	DB *sql.DB
}

func (repository TransactionRepositoryImpl) FindByTransactionId(context context.Context, tx *sql.Tx, id string) *model.TransactionModel {
	queryTransaction := "SELECT id, total_price, created_at " +
		"FROM transaction " +
		"WHERE id = $1 AND deleted_at IS NULL"
	row := tx.QueryRowContext(context, queryTransaction, id)
	transactionModel := model.TransactionModel{}
	err := row.Scan(&transactionModel.Id, &transactionModel.TotalPrice, &transactionModel.TransactionDate)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	queryDetailTransaction := "SELECT id, name, image_url, product_id, price, weight, per_unit, total, transaction_id " +
		"FROM detail_transaction " +
		"WHERE transaction_id = $1 AND deleted_at IS NULL"
	dtRows, err := tx.QueryContext(context, queryDetailTransaction, id)
	helpers.LogIfError(err)

	var detailTransactions []*model.DetailTransactionProductModel
	for dtRows.Next() {
		detailTransaction := model.DetailTransactionProductModel{}
		var imageUrl sql.NullString
		err = dtRows.Scan(&detailTransaction.Id, &detailTransaction.Name, &imageUrl, &detailTransaction.ProductId,
			&detailTransaction.Price, &detailTransaction.Weight, &detailTransaction.PerUnit, &detailTransaction.Total, &detailTransaction.TransactionId)
		if err != nil {
			continue
		}

		if imageUrl.Valid {
			detailTransaction.ImageUrl = imageUrl.String
		}
		detailTransactions = append(detailTransactions, &detailTransaction)
	}
	helpers.LogIfError(dtRows.Close())
	attachDetailTransaction(&transactionModel, detailTransactions)
	return &transactionModel
}

func attachDetailTransaction(transaction *model.TransactionModel, detailTransaction []*model.DetailTransactionProductModel) {
	for _, dt := range detailTransaction {
		if dt.TransactionId == transaction.Id {
			transaction.DetailTransaction = append(transaction.DetailTransaction, dt)
		}
	}
}

func (repository TransactionRepositoryImpl) FindByUserId(context context.Context, tx *sql.Tx, userId string) []*model.TransactionModel {
	sqlDetailTransaction := "SELECT dt.id, name, image_url, product_id, price, weight, per_unit, total, transaction_id " +
		"FROM detail_transaction dt " +
		"JOIN transaction t ON t.id = dt.transaction_id " +
		"WHERE t.user_id = $1 AND dt.deleted_at IS NULL"
	detailTransactionRows, err := tx.QueryContext(context, sqlDetailTransaction, userId)
	helpers.PanicIfError(err)
	var detailTransactions []*model.DetailTransactionProductModel
	for detailTransactionRows.Next() {
		detailTransaction := model.DetailTransactionProductModel{}
		var imageUrl sql.NullString
		err = detailTransactionRows.Scan(&detailTransaction.Id, &detailTransaction.Name, &imageUrl, &detailTransaction.ProductId,
			&detailTransaction.Price, &detailTransaction.Weight, &detailTransaction.PerUnit, &detailTransaction.Total, &detailTransaction.TransactionId)
		if err != nil {
			continue
		}

		if imageUrl.Valid {
			detailTransaction.ImageUrl = imageUrl.String
		}
		detailTransactions = append(detailTransactions, &detailTransaction)
	}
	helpers.LogIfError(detailTransactionRows.Close())

	sqlTransaction := "SELECT id, total_price, created_at " +
		"FROM transaction WHERE user_id = $1 AND deleted_at IS NULL"
	transactionRows, err := tx.QueryContext(context, sqlTransaction, userId)
	helpers.PanicIfError(err)

	var transactions []*model.TransactionModel
	for transactionRows.Next() {
		transactionModel := model.TransactionModel{}
		err = transactionRows.Scan(&transactionModel.Id, &transactionModel.TotalPrice, &transactionModel.TransactionDate)
		if err != nil {
			continue
		}
		attachDetailTransaction(&transactionModel, detailTransactions)
		transactions = append(transactions, &transactionModel)
	}
	helpers.LogIfError(transactionRows.Close())
	return transactions
}

func (repository TransactionRepositoryImpl) Save(context context.Context, tx *sql.Tx, model *model.TransactionModel) error {
	sqlTransaction := "INSERT INTO transaction(id, total_price, user_id, created_at) VALUES($1,$2,$3, NOW())"
	transactionId, _ := uuid.NewUUID()
	_, err := tx.ExecContext(context, sqlTransaction, transactionId, model.TotalPrice, model.UserId)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	var statement []string
	var values []interface{}
	for index, dt := range model.DetailTransaction {
		statement = append(statement, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			index*9+1,
			index*9+2,
			index*9+3,
			index*9+4,
			index*9+5,
			index*9+6,
			index*9+7,
			index*9+8,
			index*9+9,
			index*9+10))

		id, err := uuid.NewUUID()
		helpers.PanicIfError(err)
		now := time.Now()
		values = append(values, transactionId, dt.ProductId, id, dt.PerUnit, dt.Price, dt.Total, dt.Weight, dt.ImageUrl, dt.Name, now.Format("2006-01-02 15:04:05"))
	}

	sqlDetailTransaction := fmt.Sprintf("INSERT INTO detail_transaction(transaction_id,product_id,id,per_unit,price,total,weight,image_url,name, created_at) "+
		"VALUES %s", strings.Join(statement, ","))

	_, err = tx.ExecContext(context, sqlDetailTransaction, values...)
	helpers.LogIfError(err)
	return err
}

func (repository TransactionRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) error {
	sqlDetailTransaction := "UPDATE detail_transaction SET deleted_at = NOW() WHERE transaction_id = $1"
	_, err := tx.ExecContext(context, sqlDetailTransaction, id)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	sqlTransaction := "UPDATE transaction SET deleted_at = NOW() WHERE id = $1"
	_, err = tx.ExecContext(context, sqlTransaction, id)
	helpers.LogIfError(err)
	return err
}