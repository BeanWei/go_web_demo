package common

import (
	"time"
	"net/http"
	"github.com/labstack/echo"
	"github.com/BeanWei/go-web-demo/model"
)

type Replyinfo struct {
	ErrNo int `json:"errNo"`
	MSG string `json:"msg"`
	DATA *time.Time `json:"data"`
}

// SendErrJSON 错误发生时发送错误JSON
func (c echo.Context) SendErrJSON(msg string, args ...interface{}) error {
	if len(args) == 0 {
		panic("缺少 *echo.Context")
	}
	// var c *echo.Context
	var errNo = model.ErrorCode.ERROR
	if len(args) == 1 {
		theCtx, ok := args[0].(c *echo.Context)
		if !ok {
			panic("缺少 *echo.Context")
		}
		c = theCtx
	} else if len(args) == 2 {
		theErrNo, ok := args[0].(int)
		if !ok {
			panic("errNo不正确")
		}
		errNo = theErrNo
		theCtx, ok := args[1].(*echo.Context)
		if !ok {
			panic("缺少 *echo.Context")
		}
		c = theCtx
	}
	r := &Replyinfo{
		ErrNo: errNo,
		MSG: msg,
		DATA: time.Now()
	}
	return c.JSONPretty(http.StatusOK, r, "	")
}