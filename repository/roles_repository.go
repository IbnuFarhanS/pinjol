package repository

<<<<<<< HEAD
import "github.com/IbnuFarhanS/pinjol/model"

type RolesRepository interface {
	Save(newRoles model.Roles) (model.Roles, error)
	Update(updatedRoles model.Roles) (model.Roles, error)
	Delete(id int64) (model.Roles, error)
	FindById(id int64) (model.Roles, error)
	FindAll() ([]model.Roles, error)
	FindByName(name string) (model.Roles, error)
=======


import "github.com/IbnuFarhanS/pinjol/model"

type RoleRepository interface {
	Save(newRole model.Role) (model.Role, error)
	Update(updatedRole model.Role) (model.Role, error)
	Delete(id uint) (model.Role, error)
	FindById(id uint) (model.Role, error)
	FindAll() ([]model.Role, error)
	FindByName(name string) (model.Role, error)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
