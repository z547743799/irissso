package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"gitlab.com/z547743799/iriscommon/utils"
	"gitlab.com/z547743799/irissso/service"
)

// DeptController operate dept
//type TbUserController struct {
//	DeptDao db.TbUser
//}
//var DeptDao db.TbUser

// L.JSON(c.Writer, http.StatusOK, depts)
type TokenController struct {
	Ctx     iris.Context
	Service service.TbTokenService
}

func (c *TokenController) GetTokenBy(token string) {

	callback := c.Ctx.URLParam("callback")
	Sen := utils.Manager.Start(c.Ctx)
	jsonuser := Sen.GetString(token)

	Result := c.Service.GetUserByToken(jsonuser)

	c.Ctx.JSONP(Result, context.JSONP{Callback: callback})
	//重定向
	//	c.Ctx.Redirect("http://127.0.0.1:8082")
}
