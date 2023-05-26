package service

import "github.com/IbnuFarhanS/pinjol/data/request"

type AuthService interface {
	Login(borrower request.LoginRequest) (string, error)
	Register(borrower request.CreateUsersRequest)
}
