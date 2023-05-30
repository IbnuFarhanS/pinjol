package service

import "github.com/IbnuFarhanS/pinjol/model"

type RoleService interface {
	Save(newRole model.Role) (model.Role, error)
	Update(updatedRole model.Role) (model.Role, error)
	Delete(id uint) (model.Role, error)
	FindById(id uint) (model.Role, error)
	FindAll() ([]model.Role, error)
	FindByName(name string) (model.Role, error)
}
