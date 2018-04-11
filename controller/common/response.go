// package common

// import (
// 	"net/http"
// 	"github.com/labstack/echo"
// 	"github.com/BeanWei/go-web-demo/model"
// )

// // SendErrJSON 错误发生时发送错误JSON
// func SendErrJSON(msg string, args ...interface{}) {
// 	if len(args) == 0 {
// 		panic("缺少 *echo.Context")
// 	}
// 	var c *echo.Context
// 	var errNo = model.ErrorCode.ERROR
// 	if len(args) == 1 {
// 		theCtx, ok := args[0].(c *echo.Context)
// 		if !ok {
// 			panic("缺少 *echo.Context")
// 		}
// 		c = theCtx
// 	} else if len(args) == 2 {
// 		theErrNo, ok := args[0].(int)
// 		if !ok {
// 			panic("errNo不正确")
// 		}
// 		errNo = theErrNo
// 		theCtx, ok := args[1].(*echo.Context)
// 		if !ok {
// 			panic("缺少 *echo.Context")
// 		}
// 		c = theCtx
// 	}
// 	c.JSON(http.StatusOK, echo.)
// }