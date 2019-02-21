package controller

import (
	"github.com/kataras/iris"
	"gitlab.com/z547743799/irismanager/models"
	"gitlab.com/z547743799/irissso/service"
)

// DeptController operate dept
//type TbUserController struct {
//	DeptDao db.TbUser
//}
//var DeptDao db.TbUser

// L.JSON(c.Writer, http.StatusOK, depts)
type RegisterController struct {
	Ctx     iris.Context
	Service service.TbRegisterService
}

func (c *RegisterController) PostRegister() {
 var users  models.TbUser
	user := c.Ctx.FormValues()

		users.Username=user["username"][0]
		users.Password=user["password"][0]
		users.Phone=user["phone"][0]

	Result := c.Service.Register(users)
	c.Ctx.JSON(Result)

}
func (c *RegisterController) GetCheckBy(param string,typ int) {
	Result := c.Service.CheckData(param,typ)
	c.Ctx.JSON(Result)
}
