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
	UsersRepository repository.UsersRepository
}

func NewAuthServiceImpl(usersRepository repository.UsersRepository) AuthService {
	return &AuthServiceImpl{
		UsersRepository: usersRepository,
	}
}

// Login implements AuthenticationService
func (a *AuthServiceImpl) Login(users request.LoginRequest) (string, error) {
	// Find username in database
	new_users, users_err := a.UsersRepository.FindByUsername(users.Username)
	if users_err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_users.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_users.ID, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil

}

// Register implements AuthenticationService
func (a *AuthServiceImpl) Register(newUsers request.CreateUsersRequest) {

	hashedPassword, err := utils.HashPassword(newUsers.Password)
	helper.ErrorPanic(err)

	created_at := time.Now()

	newUser := model.Users{
		Username:     newUsers.Username,
		Password:     hashedPassword,
		Nik:          newUsers.Nik,
		Name:         newUsers.Name,
		Alamat:       newUsers.Alamat,
		Phone_Number: newUsers.Phone_Number,
		Limit:        2000000,
		Roles:        model.Roles{ID: 1},
		Created_At:   created_at,
	}
	a.UsersRepository.Save(newUser)
}

func (s *AuthServiceImpl) FindAll() ([]model.Users, error) {
	return s.UsersRepository.FindAll()
}

func (s *AuthServiceImpl) FindByUsername(username string) (model.Users, error) {
	return s.UsersRepository.FindByUsername(username)
}
