package app

import (
	"github.com/gin-gonic/gin"
	code "github.com/sixinshuier/go-web/pkg/code"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Header Header
}

type Header map[string][]string

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}, header map[string][]string) {
	g.C.JSON(httpCode, Response{
		Header: header,
		Code:   errCode,
		Msg:    code.GetMsg(errCode),
		Data:   data,
	})
	return
}
