package repository

import "github.com/IbnuFarhanS/pinjol/model"

type UserRepository interface {
	Save(newUser model.User) (model.User, error)
	Update(updatedUser model.User) (model.User, error)
	Delete(id uint) (model.User, error)
	FindById(id uint) (model.User, error)
	FindAll() ([]model.User, error)
	FindByUsername(username string) (model.User, error)
}
