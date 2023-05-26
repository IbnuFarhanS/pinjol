package repository



import (
	"errors"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

type RolesRepositoryImpl struct {
	Db *gorm.DB
}

func NewRolesRepositoryImpl(Db *gorm.DB) RolesRepository {
	return &RolesRepositoryImpl{Db: Db}
}

// Delete implements RolesRepository
func (r *RolesRepositoryImpl) Delete(id int64) (model.Roles, error) {
	var Roles model.Roles
	result := r.Db.Where("id = ?", id).Delete(&Roles)
	helper.ErrorPanic(result.Error)
	return Roles, nil
}

// FindAll implements RolesRepository
func (r *RolesRepositoryImpl) FindAll() ([]model.Roles, error) {
	var Roles []model.Roles
	results := r.Db.Find(&Roles)
	helper.ErrorPanic(results.Error)
	return Roles, nil
}

// FindById implements RolesRepository
func (r *RolesRepositoryImpl) FindById(id int64) (model.Roles, error) {
	var Roles model.Roles
	result := r.Db.Find(&Roles, id)
	if result != nil {
		return Roles, errors.New("Roles is not found")

	}
	return Roles, nil
}

// Save implements RolesRepository
func (r *RolesRepositoryImpl) Save(newRoles model.Roles) (model.Roles, error) {
	result := r.Db.Create(&newRoles)
	helper.ErrorPanic(result.Error)
	return newRoles, nil
}

// Update implements RolesRepository
func (r *RolesRepositoryImpl) Update(updatedRoles model.Roles) (model.Roles, error) {
	var rol model.Roles
	created_at := rol.Created_at

	var updateRoles = model.Roles{
		ID:         updatedRoles.ID,
		Name:       updatedRoles.Name,
		Created_at: created_at,
	}
	result := r.Db.Model(&updatedRoles).Updates(updateRoles)
	helper.ErrorPanic(result.Error)
	return updatedRoles, nil
}

// FindByUsername implements RolesRepository
func (r *RolesRepositoryImpl) FindByUsername(username string) (model.Roles, error) {
	var Roles model.Roles
	result := r.Db.First(&Roles, "username = ?", username)

	if result.Error != nil {
		return Roles, errors.New("invalid username or Password")
	}
	return Roles, nil
}
