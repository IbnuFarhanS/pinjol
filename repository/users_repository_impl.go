package repository

import (
	"errors"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

// Delete implements UsersRepository
func (u *UsersRepositoryImpl) Delete(id int64) (model.Users, error) {
	var users model.Users
	result := u.Db.Where("id = ?", id).Delete(&users)
	helper.ErrorPanic(result.Error)
	return users, nil
}

// FindAll implements UsersRepository
func (u *UsersRepositoryImpl) FindAll() ([]model.Users, error) {
	var users []model.Users
	results := u.Db.Find(&users)
	helper.ErrorPanic(results.Error)
	return users, nil
}

// FindById implements UsersRepository
func (u *UsersRepositoryImpl) FindById(id int64) (model.Users, error) {
	var users model.Users
	result := u.Db.Find(&users, id)
	if result != nil {
		return users, errors.New("users is not found")
	}
	return users, nil
}

// Save implements UsersRepository
func (u *UsersRepositoryImpl) Save(newUsers model.Users) (model.Users, error) {
	result := u.Db.Create(&newUsers)
	helper.ErrorPanic(result.Error)
	return newUsers, nil
}

// Update implements UsersRepository
func (u *UsersRepositoryImpl) Update(updatedUsers model.Users) (model.Users, error) {
	var users model.Users
	updatedUsers.Created_At = users.Created_At
	updatedUsers.RolesID = users.RolesID
	var updateUsers = model.Users{
		ID:           updatedUsers.ID,
		Username:     updatedUsers.Username,
		Password:     updatedUsers.Password,
		Nik:          updatedUsers.Nik,
		Name:         updatedUsers.Name,
		Alamat:       updatedUsers.Alamat,
		Phone_Number: updatedUsers.Phone_Number,
		Limit:        updatedUsers.Limit,
		RolesID:      updatedUsers.RolesID,
		Created_At:   updatedUsers.Created_At,
	}
	result := u.Db.Model(&updatedUsers).Updates(updateUsers)
	helper.ErrorPanic(result.Error)
	return updatedUsers, nil
}

// FindByUsername implements UsersRepository
func (u *UsersRepositoryImpl) FindByUsername(username string) (model.Users, error) {
	var users model.Users
	result := u.Db.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("invalid username or Password")
	}
	return users, nil
}
