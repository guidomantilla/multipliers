package endpoints

import (
	"github.com/gin-gonic/gin"
)

type NumbersEndpoint interface {
	Save(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}
