package repository

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"github.com/ramdanariadi/grocery-product-service/main/product/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"log"
)

type ProductRepositoryImpl struct{}

func (repository ProductRepositoryImpl) FindById(context context.Context, tx *sql.Tx, id string) *model.ProductModel {
	query := "SELECT products.id, name, price, per_unit, weight, categories.category, category_id, description, products.image_url  " +
		"FROM products " +
		"JOIN categories ON products.category_id = categories.id " +
		"WHERE products.id = $1 AND products.deleted_at IS NULL"
	row := tx.QueryRowContext(context, query, id)
	product := model.ProductModel{}
	var imageUrl sql.NullString
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.PerUnit, &product.Weight,
		&product.Category, &product.CategoryId, &product.Description, &imageUrl)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if imageUrl.Valid {
		product.ImageUrl = imageUrl.String
	}

	return &product
}

func (repository ProductRepositoryImpl) FindByIds(context context.Context, tx *sql.Tx, ids []string) []*model.ProductModel {
	query := "SELECT products.id, name, price, per_unit, weight, categories.category, category_id, description, products.image_url " +
		"FROM products " +
		"JOIN categories ON products.category_id = categories.id " +
		"WHERE products.id = ANY($1) AND products.deleted_at IS NULL"
	rows, err := tx.QueryContext(context, query, pq.Array(ids))
	utils.LogIfError(err)
	var products []*model.ProductModel

	for rows.Next() {
		product := model.ProductModel{}
		var imageUrl sql.NullString
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.PerUnit, &product.Weight,
			&product.Category, &product.CategoryId, &product.Description, &imageUrl)

		if err != nil {
			log.Printf("error : %s", err.Error())
			continue
		}

		if imageUrl.Valid {
			product.ImageUrl = imageUrl.String
		}
		products = append(products, &product)
	}
	return products
}

func (repository ProductRepositoryImpl) FindAll(context context.Context, tx *sql.Tx) *sql.Rows {
	query := "SELECT products.id, name, price, per_unit, weight, categories.category, category_id, description, products.image_url  " +
		"FROM products " +
		"JOIN categories ON products.category_id = categories.id " +
		"WHERE products.deleted_at IS NULL"

	rows, err := tx.QueryContext(context, query)
	utils.LogIfError(err)
	return rows
}

func (repository ProductRepositoryImpl) FindByCategory(context context.Context, tx *sql.Tx, id string) *sql.Rows {
	query := "SELECT products.id, name, price, per_unit, weight, categories.category, category_id, description, products.image_url  " +
		"FROM products " +
		"JOIN categories ON products.category_id = categories.id " +
		"WHERE products.category_id = $1 AND products.deleted_at IS NULL"
	rows, err := tx.QueryContext(context, query, id)
	utils.LogIfError(err)
	return rows
}

func (repository ProductRepositoryImpl) FindWhere(context context.Context, tx *sql.Tx, where string, value ...any) *sql.Rows {
	query := "SELECT products.id, name, price, per_unit, weight, categories.category, category_id, description, products.image_url  " +
		"FROM products " +
		"JOIN categories ON products.category_id = categories.id " +
		"WHERE " + where
	rows, err := tx.QueryContext(context, query, value...)
	utils.LogIfError(err)
	return rows
}

func (repository ProductRepositoryImpl) Save(context context.Context, tx *sql.Tx, product *model.ProductModel) error {
	sql := "INSERT INTO products(id, name, weight, price, per_unit, category_id, description, image_url, is_top, is_recommended, created_at) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,NOW())"
	_, err := tx.ExecContext(context, sql, product.Id, product.Name, product.Weight, product.Price,
		product.PerUnit, product.CategoryId, product.Description, product.ImageUrl, false, false)
	utils.LogIfError(err)
	return err
}

func (repository ProductRepositoryImpl) Update(context context.Context, tx *sql.Tx, product *model.ProductModel) error {
	sql := "UPDATE products SET name=$1, price=$2, weight=$3, category_id=$4, per_unit=$5," +
		"description=$6, image_url=$7, is_top=$8, is_recommended=$9, updated_at = NOW()" +
		"WHERE id = $10 AND deleted_at IS NULL"
	_, err := tx.ExecContext(context, sql, product.Name, product.Price, product.Weight,
		product.CategoryId, product.PerUnit, product.Description, product.ImageUrl, product.IsTop,
		product.IsRecommended, product.Id)
	utils.LogIfError(err)
	return err
}

func (repository ProductRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) error {
	sql := "UPDATE products set deleted_at = NOW() WHERE id = $1"
	_, err := tx.ExecContext(context, sql, id)
	utils.LogIfError(err)
	return err
}
