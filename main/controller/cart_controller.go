package controller

import "github.com/gin-gonic/gin"

type CartController interface {
	Store(ctx *gin.Context)
	Destroy(ctx *gin.Context)
	Find(ctx *gin.Context)
}
