package controller

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"gitlab.com/z547743799/iriscommon/utils"
	"gitlab.com/z547743799/irissso/service"
)

// DeptController operate dept
//type TbUserController struct {
//	DeptDao db.TbUser
//}
//var DeptDao db.TbUser

// L.JSON(c.Writer, http.StatusOK, depts)
type LoginController struct {
	Ctx iris.Context

	Service service.TbLoginService
}

func (c *LoginController) PostLogin() {
	username := c.Ctx.PostValue("username")
	password := c.Ctx.PostValue("password")
	Result, data := c.Service.UserLogin(username, password)
	if data != nil {

		Sen := utils.Manager.Start(c.Ctx)
		cookie := sessions.GetCookie(c.Ctx, "token")
		Sen.Set(cookie, string(data))
		jsonuser := Sen.GetString(cookie)
		fmt.Println(jsonuser)
		Result.Data = cookie
	}

	c.Ctx.JSON(Result)
}
