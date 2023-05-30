package service

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/config"
	"github.com/IbnuFarhanS/pinjol/data/request"
	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/IbnuFarhanS/pinjol/utils"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
}

func NewAuthServiceImpl(UserRepository repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		UserRepository: UserRepository,
	}
}

// Login implements AuthenticationService
func (a *AuthServiceImpl) Login(User request.LoginRequest) (string, error) {
	// Find username in database
	new_User, User_err := a.UserRepository.FindByUsername(User.Username)
	if User_err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_User.Password, User.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_User.ID, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil

}

// Register implements AuthenticationService
func (a *AuthServiceImpl) Register(newUser request.CreateUsersRequest) {

	hashedPassword, err := utils.HashPassword(newUser.Password)
	helper.ErrorPanic(err)

	created_at := time.Now()

	id := model.Role{ID: 1}

	newUsers := model.User{
		Username:    newUser.Username,
		Password:    hashedPassword,
		NIK:         newUser.NIK,
		Name:        newUser.Name,
		Address:     newUser.Address,
		PhoneNumber: newUser.PhoneNumber,
		Limit:       2000000,
		Role:        id,
		CreatedAt:   created_at,
	}
	a.UserRepository.Save(newUsers)
}

func (s *AuthServiceImpl) FindAll() ([]model.User, error) {
	return s.UserRepository.FindAll()
}

func (s *AuthServiceImpl) FindByUsername(username string) (model.User, error) {
	return s.UserRepository.FindByUsername(username)
}

func (s *AuthServiceImpl) FindByUserID(id uint) (model.User, error) {
	return s.UserRepository.FindById(id)
}
