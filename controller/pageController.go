package controller

import (
	"github.com/kataras/iris"
)

// DeptController operate dept
//type TbUserController struct {
//	DeptDao db.TbUser
//}
//var DeptDao db.TbUser

// L.JSON(c.Writer, http.StatusOK, depts)
type Controller struct {
	Ctx iris.Context
}

func (c *Controller) GetLogin() {

	c.Ctx.View("login.html")

}

func (c *Controller) GetRegister() {
	c.Ctx.View("register.html")
}
