package service

import (
	"crypto/md5"
	"fmt"
	"time"

	"gitlab.com/z547743799/irismanager/db"
	"gitlab.com/z547743799/irismanager/models"

	"github.com/go-xorm/xorm"
	"gitlab.com/z547743799/iriscommon/utils"
)

type TbRegisterService interface {
	Register(user models.TbUser) *utils.E3Result
	CheckData(param string, typ int) *utils.E3Result
}

type tbRegisterService struct {
	engine *xorm.Engine
}

func NewTbRegisterService() TbRegisterService {
	return &tbRegisterService{
		engine: db.X,
	}
}

func (d *tbRegisterService) Register(user models.TbUser) *utils.E3Result {
	var (
		result *utils.E3Result
		ok     bool
	)

	result = d.CheckData(user.Username, 1)
	ok = result.Data.(bool)
	if !ok {
		return utils.Build(400, "用户名以存在")
	}

	result = d.CheckData(user.Phone, 2)
	ok = result.Data.(bool)
	if !ok {
		return utils.Build(400, "电话号码以存在")
	}
	user.Created = time.Now()
	user.Updated = time.Now()

	md := md5.New()
	md.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x",md.Sum(nil))

	//d.engine.SQL(`insert into tb_user (username,password,phone)values(?,?,?)`,user.Username,user.Password,user.Phone)
	_, err := d.engine.Insert(&user)
	if err != nil {
		panic(err)
	}
	return utils.Ok(nil)

}

func (d *tbRegisterService) CheckData(param string, typ int) *utils.E3Result {

	tbuser := new(models.TbUser)
	switch typ {
	case 1:
		_, err := d.engine.Where("username=?", param).Desc("id").Get(tbuser)
		if err != nil || tbuser.Username != "" {
			return utils.Ok(false)
		}
	case 2:
		_, err := d.engine.Where("phone=?", param).Desc("id").Get(tbuser)
		if err != nil || tbuser.Phone != "" {
			return utils.Ok(false)
		}
	}

	return utils.Ok(true)
}
