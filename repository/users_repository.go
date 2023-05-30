package repository

import "github.com/IbnuFarhanS/pinjol/model"

<<<<<<< HEAD
type UsersRepository interface {
	Save(newUsers model.Users) (model.Users, error)
	Update(updatedUsers model.Users) (model.Users, error)
	Delete(id int64) (model.Users, error)
	FindById(id int64) (model.Users, error)
	FindAll() ([]model.Users, error)
	FindByUsername(username string) (model.Users, error)
=======
type UserRepository interface {
	Save(newUser model.User) (model.User, error)
	Update(updatedUser model.User) (model.User, error)
	Delete(id uint) (model.User, error)
	FindById(id uint) (model.User, error)
	FindAll() ([]model.User, error)
	FindByUsername(username string) (model.User, error)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
