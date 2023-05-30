package service

import (
	"github.com/IbnuFarhanS/pinjol/data/request"
	"github.com/IbnuFarhanS/pinjol/model"
)

type AuthService interface {
<<<<<<< HEAD
	Login(users request.LoginRequest) (string, error)
	Register(newUsers request.CreateUsersRequest)
	FindAll() ([]model.Users, error)
	FindByUsername(username string) (model.Users, error)
=======
	Login(User request.LoginRequest) (string, error)
	Register(newUser request.CreateUsersRequest)
	FindAll() ([]model.User, error)
	FindByUsername(username string) (model.User, error)
	FindByUserID(id uint) (model.User, error)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
