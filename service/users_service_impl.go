package service

import (
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	reposity "github.com/IbnuFarhanS/pinjol/repository"
	"github.com/IbnuFarhanS/pinjol/utils"
	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UsersRepository reposity.UsersRepository
	Validate        *validator.Validate
}

// Delete implements UsersService
func (s *UsersServiceImpl) Delete(id int64) (model.Users, error) {
	return s.UsersRepository.Delete(id)
}

// FindAll implements UsersService
func (s *UsersServiceImpl) FindAll() ([]model.Users, error) {
	return s.UsersRepository.FindAll()
}

// FindById implements UsersService
func (s *UsersServiceImpl) FindById(id int64) (model.Users, error) {
	return s.UsersRepository.FindById(id)
}

// FindByUsername implements UsersService
func (s *UsersServiceImpl) FindByUsername(username string) (model.Users, error) {
	return s.UsersRepository.FindByUsername(username)
}

// Save implements UsersService
func (s *UsersServiceImpl) Save(newUsers model.Users) (model.Users, error) {
	hashedPassword, err := utils.HashPassword(newUsers.Password)
	helper.ErrorPanic(err)

	newUser := model.Users{
		Username:     newUsers.Username,
		Password:     hashedPassword,
		Nik:          newUsers.Nik,
		Name:         newUsers.Name,
		Alamat:       newUsers.Alamat,
		Phone_Number: newUsers.Phone_Number,
		Limit:        2000000,
		Roles:        model.Roles{ID: 1},
		Created_At:   time.Now(),
	}
	s.UsersRepository.Save(newUser)
	return newUser, nil
}

// Update implements UsersService
func (s *UsersServiceImpl) Update(updatedUsers model.Users) (model.Users, error) {
	hashedPassword, err := utils.HashPassword(updatedUsers.Password)
	helper.ErrorPanic(err)

	var bor model.Users
	create_at := bor.Created_At

	newUser := model.Users{
		ID:           updatedUsers.ID,
		Username:     updatedUsers.Username,
		Password:     hashedPassword,
		Nik:          updatedUsers.Nik,
		Name:         updatedUsers.Name,
		Alamat:       updatedUsers.Alamat,
		Phone_Number: updatedUsers.Phone_Number,
		Limit:        updatedUsers.Limit,
		Created_At:   create_at,
	}

	return s.UsersRepository.Update(newUser)
}

func NewUsersServiceImpl(UsersRepository reposity.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: UsersRepository,
		Validate:        validate,
	}
}
