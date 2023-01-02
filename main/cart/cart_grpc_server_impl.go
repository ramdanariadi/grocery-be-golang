package cart

import (
	"context"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/cart/model"
	productModel "github.com/ramdanariadi/grocery-product-service/main/product/model"
	"github.com/ramdanariadi/grocery-product-service/main/response"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
)

type CartServiceServerImpl struct {
	DB *gorm.DB
}

func NewCartServiceImpl(db *gorm.DB) *CartServiceServerImpl {
	return &CartServiceServerImpl{
		DB: db,
	}
}

func (server CartServiceServerImpl) Save(_ context.Context, cart *Cart) (*response.Response, error) {
	var productRef productModel.Product
	first := server.DB.First(&productRef, "id = ?", cart.ProductId)
	if first.RowsAffected == 0 {
		status, message := utils.QueryResponse(false)
		return &response.Response{
			Message: message,
			Status:  status,
		}, nil
	}

	var cartModel model.Cart
	tx := server.DB.First(&cartModel, "product_id = ? and user_id = ?", cart.ProductId, cart.UserId)
	if tx.RowsAffected > 0 {
		cartModel.Total = cartModel.Total + cart.Total
	} else {
		id, _ := uuid.NewUUID()
		cartModel = model.Cart{
			ID:        id.String(),
			ImageUrl:  productRef.ImageUrl,
			ProductId: productRef.ID,
			Name:      productRef.Name,
			Weight:    uint32(productRef.Weight),
			Category:  productRef.Category,
			Price:     productRef.Price,
			PerUnit:   uint64(productRef.PerUnit),
			UserId:    cart.UserId,
			Total:     cart.Total,
		}
	}
	save := server.DB.Save(&cartModel)
	status, message := utils.ModifyingResponse(save.Error == nil)
	return &response.Response{
		Message: message,
		Status:  status,
	}, nil
}

func (server CartServiceServerImpl) Delete(_ context.Context, id *CartAndUserId) (*response.Response, error) {
	tx := server.DB.Delete(&model.Cart{ID: id.Id, UserId: id.UserId})
	status, message := utils.ModifyingResponse(tx.Error == nil)
	return &response.Response{Message: message, Status: status}, nil
}

func (server CartServiceServerImpl) FindByUserId(_ context.Context, userId *CartUserId) (*MultipleCartResponse, error) {
	var carts []*model.Cart
	tx := server.DB.Find(&carts, "user_id = ?", userId.Id)
	utils.LogIfError(tx.Error)
	wishlist := fetchCarts(carts)
	status, message := utils.QueryResponse(true)
	return &MultipleCartResponse{Status: status, Message: message, Data: wishlist}, nil
}

func fetchCarts(carts []*model.Cart) []*CartDetail {
	var cartDetails []*CartDetail
	for _, cart := range carts {
		cartDetail := CartDetail{
			Id:        cart.ID,
			Name:      cart.Name,
			Price:     cart.Price,
			Weight:    cart.Weight,
			Category:  cart.Category,
			Total:     cart.Total,
			PerUnit:   uint32(cart.PerUnit),
			ImageUrl:  cart.ImageUrl,
			ProductId: cart.ProductId,
			UserId:    cart.UserId,
		}
		cartDetails = append(cartDetails, &cartDetail)
	}
	return cartDetails
}

func (server CartServiceServerImpl) mustEmbedUnimplementedCartServiceServer() {
	//TODO implement me
	panic("implement me")
}
