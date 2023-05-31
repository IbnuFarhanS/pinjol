package service

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/IbnuFarhanS/pinjol/utils"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

// Delete implements UserService
func (s *UserServiceImpl) Delete(id uint) (model.User, error) {
	return s.UserRepository.Delete(id)
}

// FindAll implements UserService
func (s *UserServiceImpl) FindAll() ([]model.User, error) {
	return s.UserRepository.FindAll()
}

// FindById implements UserService
func (s *UserServiceImpl) FindById(id uint) (model.User, error) {
	return s.UserRepository.FindById(id)
}

// FindByUsername implements UserService
func (s *UserServiceImpl) FindByUsername(username string) (model.User, error) {
	return s.UserRepository.FindByUsername(username)
}

// Save implements UserService
func (s *UserServiceImpl) Save(newUser model.User) (model.User, error) {
	if newUser.Username == "" {
		return model.User{}, errors.New("username is required")
	}
	// Check if username already exists
	existingUser, err := s.UserRepository.FindByUsername(newUser.Username)
	if err != nil {
		return model.User{}, err
	}
	if existingUser.ID != 0 {
		return model.User{}, errors.New("username is already in use")
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	helper.ErrorPanic(err)

	newUsers := model.User{
		Username:    newUser.Username,
		Password:    hashedPassword,
		NIK:         newUser.NIK,
		Name:        newUser.Name,
		Address:     newUser.Address,
		PhoneNumber: newUser.PhoneNumber,
		Limit:       2000000,
		Role:        model.Role{ID: 1},
		CreatedAt:   time.Now(),
	}
	s.UserRepository.Save(newUsers)
	return newUser, nil
}

// Update implements UserService
func (s *UserServiceImpl) Update(updatedUser model.User) (model.User, error) {
	hashedPassword, err := utils.HashPassword(updatedUser.Password)
	helper.ErrorPanic(err)

	var bor model.User
	create_at := bor.CreatedAt

	newUser := model.User{
		ID:          updatedUser.ID,
		Username:    updatedUser.Username,
		Password:    hashedPassword,
		NIK:         updatedUser.NIK,
		Name:        updatedUser.Name,
		Address:     updatedUser.Address,
		PhoneNumber: updatedUser.PhoneNumber,
		Limit:       updatedUser.Limit,
		CreatedAt:   create_at,
	}

	return s.UserRepository.Update(newUser)
}

func NewUserServiceImpl(UserRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: UserRepository,
	}
}
