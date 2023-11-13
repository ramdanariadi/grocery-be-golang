package controller

import "github.com/gin-gonic/gin"

type TransactionController interface {
	Save(ctx *gin.Context)
	Find(ctx *gin.Context)
}
