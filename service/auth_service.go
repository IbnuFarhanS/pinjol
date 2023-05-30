package service

import (
	"github.com/IbnuFarhanS/pinjol/data/request"
	"github.com/IbnuFarhanS/pinjol/model"
)

type AuthService interface {
	Login(users request.LoginRequest) (string, error)
	Register(newUsers request.CreateUsersRequest)
	FindAll() ([]model.Users, error)
	FindByUsername(username string) (model.Users, error)
	FindByUserID(id int64) (model.Users, error)
}
