package app

import (
	"github.com/gin-gonic/gin"

	"github.com/iamlockon/gorestemplate/pkg/error"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"mgs":  error.GetMsg(errCode),
		"data": data,
	})

	return
}
