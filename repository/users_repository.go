package reposity

import "github.com/IbnuFarhanS/pinjol/model"

type UsersRepository interface {
	Save(newUsers model.Users) (model.Users, error)
	Update(updatedUsers model.Users) (model.Users, error)
	Delete(id int64) (model.Users, error)
	FindById(id int64) (model.Users, error)
	FindAll() ([]model.Users, error)
	FindByUsername(username string) (model.Users, error)
}
