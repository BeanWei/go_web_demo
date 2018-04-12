// package user

// import (
// 	"fmt"

// 	"github.com/BeanWei/go-web-demo/utils"
// 	"github.com/BeanWei/go-web-demo/config"
// 	"github.com/BeanWei/go-web-demo/model"
// 	"github.com/BeanWei/go-web-demo/controller/common"

// 	"github.com/garyburd/redigo/redis"
// 	"github.com/labstack/echo"
// )

// type Login struct {
// 	Account string `json:"account" form:"account" query:"account"`
// 	Password string `json:"password" form:"password" query:"password"`
// }

// // Login 用户登录
// func Login(c echo.Context) {
// 	SendErrJSON := common.SendErrJSON
// 	lg := new(Login)
// 	if err = c.Bind(lg); err != nil {
// 		SendErrJSON("账号或密码错误", lg)
// 		return
// 	}
// 	var user model.User
// 	if err := model.engine.Cols(lg.Account).Get(&user); err != nil {
// 		SendErrJSON("账号不存在", lg)
// 		return
// 	}
// 	if user.CheckPassword(lg.Password) {
// 		if
// 	}

// }