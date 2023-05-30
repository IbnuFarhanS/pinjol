package repository

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Delete implements UserRepository
func (u *UserRepositoryImpl) Delete(id uint) (model.User, error) {
	var User model.User
	result := u.Db.Where("id = ?", id).Delete(&User)
	helper.ErrorPanic(result.Error)
	return User, nil
}

// FindAll implements UserRepository
func (u *UserRepositoryImpl) FindAll() ([]model.User, error) {
	var User []model.User
	results := u.Db.Find(&User)
	helper.ErrorPanic(results.Error)
	return User, nil
}

// FindById implements UserRepository
func (u *UserRepositoryImpl) FindById(id uint) (model.User, error) {
	var User model.User
	result := u.Db.First(&User, id)
	if result.Error != nil {
		return User, errors.New("user not found")
	}
	return User, nil
}

// Save implements UserRepository
func (u *UserRepositoryImpl) Save(newUser model.User) (model.User, error) {
	result := u.Db.Create(&newUser)
	if result.Error != nil {
		return newUser, errors.New("user not found")
	}
	return newUser, nil
}

// Update implements UserRepository
func (u *UserRepositoryImpl) Update(updatedUser model.User) (model.User, error) {
	var User model.User
	updatedUser.CreatedAt = User.CreatedAt
	updatedUser.RoleID = User.RoleID
	var updateUser = model.User{
		ID:          updatedUser.ID,
		Username:    updatedUser.Username,
		Password:    updatedUser.Password,
		NIK:         updatedUser.NIK,
		Name:        updatedUser.Name,
		Address:     updatedUser.Address,
		PhoneNumber: updatedUser.PhoneNumber,
		Limit:       updatedUser.Limit,
		RoleID:      0,
		CreatedAt:   time.Time{},
		Role:        model.Role{},
	}
	result := u.Db.Model(&updatedUser).Updates(updateUser)
	helper.ErrorPanic(result.Error)
	return updatedUser, nil
}

// FindByUsername implements UserRepository
func (u *UserRepositoryImpl) FindByUsername(username string) (model.User, error) {
	var User model.User
	result := u.Db.First(&User, "username = ?", username)

	if result.Error != nil {
		return User, errors.New("invalid username or Password")
	}
	return User, nil
}
