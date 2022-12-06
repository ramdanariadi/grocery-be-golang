package cart

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/models"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/product"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/transactions"
	"github.com/ramdanariadi/grocery-product-service/main/service/response"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
)

type CartServiceServerImpl struct {
	Repository        transactions.CartRepositoryImpl
	ProductRepository product.ProductRepositoryImpl
}

func NewCartServiceImpl(db *sql.DB) *CartServiceServerImpl {
	return &CartServiceServerImpl{
		Repository:        transactions.CartRepositoryImpl{DB: db},
		ProductRepository: product.ProductRepositoryImpl{DB: db},
	}
}

func (server CartServiceServerImpl) Save(ctx context.Context, cart *Cart) (*response.Response, error) {
	tx, _ := server.ProductRepository.DB.Begin()
	productModel := server.ProductRepository.FindById(ctx, tx, cart.ProductId)
	cartModel := models.CartModel{
		ImageUrl:  productModel.ImageUrl,
		ProductId: productModel.Id,
		Name:      productModel.Name,
		Weight:    productModel.Weight,
		Category:  productModel.Category,
		Price:     productModel.Price,
		PerUnit:   productModel.PerUnit,
		UserId:    cart.UserId,
		Total:     cart.Total,
	}
	saved := server.Repository.Save(ctx, tx, cartModel)
	status, message := utils.FetchResponseForModifying(saved)
	return &response.Response{
		Message: message,
		Status:  status,
	}, nil
}

func (server CartServiceServerImpl) Delete(ctx context.Context, id *CartAndUserId) (*response.Response, error) {
	tx, _ := server.Repository.DB.Begin()
	deleted := server.Repository.Delete(ctx, tx, id.UserId, id.Id)
	status, message := utils.FetchResponseForModifying(deleted)
	return &response.Response{Message: message, Status: status}, nil
}

func (server CartServiceServerImpl) FindByUserId(ctx context.Context, id *CartUserId) (*MultipleCartResponse, error) {
	tx, _ := server.Repository.DB.Begin()
	rows := server.Repository.FindByUserId(ctx, tx, id.Id)
	wishlist := fetchWishlist(rows)
	satus, message := utils.FetchResponseForQuerying(len(wishlist) > 0)
	return &MultipleCartResponse{Status: satus, Message: message, Data: wishlist}, nil
}

func fetchWishlist(rows *sql.Rows) []*CartDetail {
	var carts []*CartDetail
	for rows.Next() {
		cart := CartDetail{}
		rows.Scan(&cart.Id, &cart.Name, &cart.Price, &cart.Weight, &cart.Category, &cart.Total, &cart.PerUnit, &cart.ImageUrl)
		carts = append(carts, &cart)
	}
	return carts
}

func (server CartServiceServerImpl) mustEmbedUnimplementedCartServiceServer() {
	//TODO implement me
	panic("implement me")
}