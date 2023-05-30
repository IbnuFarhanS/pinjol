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
<<<<<<< HEAD
	UsersRepository repository.UsersRepository
}

func NewAuthServiceImpl(usersRepository repository.UsersRepository) AuthService {
	return &AuthServiceImpl{
		UsersRepository: usersRepository,
=======
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
}

func NewAuthServiceImpl(UserRepository repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		UserRepository: UserRepository,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}
}

// Login implements AuthenticationService
<<<<<<< HEAD
func (a *AuthServiceImpl) Login(users request.LoginRequest) (string, error) {
	// Find username in database
	new_users, users_err := a.UsersRepository.FindByUsername(users.Username)
	if users_err != nil {
=======
func (a *AuthServiceImpl) Login(User request.LoginRequest) (string, error) {
	// Find username in database
	new_User, User_err := a.UserRepository.FindByUsername(User.Username)
	if User_err != nil {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
		return "", errors.New("invalid username or Password")
	}

	config, _ := config.LoadConfig(".")

<<<<<<< HEAD
	verify_error := utils.VerifyPassword(new_users.Password, users.Password)
=======
	verify_error := utils.VerifyPassword(new_User.Password, User.Password)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
<<<<<<< HEAD
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_users.ID, config.TokenSecret)
=======
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_User.ID, config.TokenSecret)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	helper.ErrorPanic(err_token)
	return token, nil

}

// Register implements AuthenticationService
<<<<<<< HEAD
func (a *AuthServiceImpl) Register(newUsers request.CreateUsersRequest) {

	hashedPassword, err := utils.HashPassword(newUsers.Password)
=======
func (a *AuthServiceImpl) Register(newUser request.CreateUsersRequest) {

	hashedPassword, err := utils.HashPassword(newUser.Password)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	helper.ErrorPanic(err)

	created_at := time.Now()

<<<<<<< HEAD
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
=======
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
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
