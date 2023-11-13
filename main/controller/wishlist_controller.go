package controller

import "github.com/gin-gonic/gin"

type WishlistController interface {
	Store(ctx *gin.Context)
	Destroy(ctx *gin.Context)
	Find(ctx *gin.Context)
	FindByProductId(ctx *gin.Context)
}
