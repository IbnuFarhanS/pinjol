package service

import (
	"github.com/IbnuFarhanS/pinjol/data/request"
	"github.com/IbnuFarhanS/pinjol/model"
)

type AuthService interface {
	Login(User request.LoginRequest) (string, error)
	Register(newUser request.CreateUsersRequest)
	FindAll() ([]model.User, error)
	FindByUsername(username string) (model.User, error)
	FindByUserID(id uint) (model.User, error)
}
