package service

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/go-xorm/xorm"
	"gitlab.com/z547743799/iriscommon/utils"
	"gitlab.com/z547743799/irismanager/db"
	"gitlab.com/z547743799/irismanager/models"
)

type TbLoginService interface {
	UserLogin(username, password string) (*utils.E3Result, []byte)
}

type tbLoginService struct {
	engine *xorm.Engine
}

func NewTbLoginService() TbLoginService {
	return &tbLoginService{
		engine: db.X,
	}
}

func (d *tbLoginService) UserLogin(username, password string) (*utils.E3Result, []byte) {

	user := models.TbUser{}
	md := md5.New()
	md.Write([]byte(password))
	password = fmt.Sprintf("%x", md.Sum(nil))

	_, err := d.engine.Where("username=? and password=? ", username, password).Desc("id").Get(&user)

	if err != nil {
		return utils.Build(400, "用户名或密码错误"), nil
	}

	//pool := rediscli.NewRedisPool(redisinit.Url).Get()
	//defer pool.Close()
	user.Password = ""
	data, err := json.Marshal(user)
	if err != nil {
		utils.Build(400, fmt.Sprintln(err))
	}

	//pool.Do("set", token, data, "EX", 60*time.Second)
	if err != nil {
		utils.Build(400, fmt.Sprintln(err))
	}

	return utils.Ok(nil), data

}
