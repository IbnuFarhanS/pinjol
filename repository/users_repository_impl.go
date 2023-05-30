package repository

import (
	"errors"
<<<<<<< HEAD
=======
	"time"
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

<<<<<<< HEAD
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
	result := u.Db.First(&users, id)
	if result.Error != nil {
		return users, errors.New("user not found")
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
=======
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
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
