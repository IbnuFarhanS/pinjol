package repository

import (
	"errors"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	Db *gorm.DB
}

func NewRoleRepositoryImpl(Db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{Db: Db}
}

// Delete implements RoleRepository
func (r *RoleRepositoryImpl) Delete(id uint) (model.Role, error) {
	var rol model.Role
	result := r.Db.Where("id = ?", id).Delete(&rol)
	helper.ErrorPanic(result.Error)
	return rol, nil
}

// FindAll implements RoleRepository
func (r *RoleRepositoryImpl) FindAll() ([]model.Role, error) {
	var rol []model.Role
	results := r.Db.Find(&rol)
	helper.ErrorPanic(results.Error)
	return rol, nil
}

// FindById implements RoleRepository
func (r *RoleRepositoryImpl) FindById(id uint) (model.Role, error) {
	var rol model.Role
	result := r.Db.Find(&rol, id)
	if result != nil {
		return rol, errors.New("role is not found")

	}
	return rol, nil
}

// Save implements RoleRepository
func (r *RoleRepositoryImpl) Save(newRole model.Role) (model.Role, error) {
	result := r.Db.Create(&newRole)
	helper.ErrorPanic(result.Error)
	return newRole, nil
}

// Update implements RoleRepository
func (r *RoleRepositoryImpl) Update(updatedRole model.Role) (model.Role, error) {
	var rol model.Role
	created_at := rol.CreatedAt

	var updateRole = model.Role{
		ID:        updatedRole.ID,
		Name:      updatedRole.Name,
		CreatedAt: created_at,
	}
	result := r.Db.Model(&updatedRole).Updates(updateRole)
	helper.ErrorPanic(result.Error)
	return updatedRole, nil
}

// FindByUsername implements RoleRepository
func (r *RoleRepositoryImpl) FindByName(name string) (model.Role, error) {
	var rol model.Role
	result := r.Db.First(&rol, "name = ?", name)

	if result.Error != nil {
		return rol, errors.New("invalid name")
	}
	return rol, nil
}
