package service

import (
	"encoding/json"
	"github.com/go-xorm/xorm"
	"gitlab.com/z547743799/iriscommon/utils"
	"gitlab.com/z547743799/irismanager/db"
	"gitlab.com/z547743799/irismanager/models"
)

type TbTokenService interface {
	GetUserByToken(token string) *utils.E3Result
}

type tbTokenService struct {
	engine *xorm.Engine
}

func NewTokenService() TbTokenService {
	return &tbTokenService{
		engine: db.X,
	}
}

func (d *tbTokenService) GetUserByToken(jsonuser string) *utils.E3Result {

	User:=models.TbUser{}
	json.Unmarshal([]byte(jsonuser),&User)

	return utils.Ok(User)

}
