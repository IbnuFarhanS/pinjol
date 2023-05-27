package repository



import "github.com/IbnuFarhanS/pinjol/model"

type RolesRepository interface {
	Save(newRoles model.Roles) (model.Roles, error)
	Update(updatedRoles model.Roles) (model.Roles, error)
	Delete(id int64) (model.Roles, error)
	FindById(id int64) (model.Roles, error)
	FindAll() ([]model.Roles, error)
	FindByName(name string) (model.Roles, error)
}
