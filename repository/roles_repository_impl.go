package repository

import (
	"errors"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

<<<<<<< HEAD
type RolesRepositoryImpl struct {
	Db *gorm.DB
}

func NewRolesRepositoryImpl(Db *gorm.DB) RolesRepository {
	return &RolesRepositoryImpl{Db: Db}
}

// Delete implements RolesRepository
func (r *RolesRepositoryImpl) Delete(id int64) (model.Roles, error) {
	var rol model.Roles
=======
type RoleRepositoryImpl struct {
	Db *gorm.DB
}

func NewRoleRepositoryImpl(Db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{Db: Db}
}

// Delete implements RoleRepository
func (r *RoleRepositoryImpl) Delete(id uint) (model.Role, error) {
	var rol model.Role
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	result := r.Db.Where("id = ?", id).Delete(&rol)
	helper.ErrorPanic(result.Error)
	return rol, nil
}

<<<<<<< HEAD
// FindAll implements RolesRepository
func (r *RolesRepositoryImpl) FindAll() ([]model.Roles, error) {
	var rol []model.Roles
=======
// FindAll implements RoleRepository
func (r *RoleRepositoryImpl) FindAll() ([]model.Role, error) {
	var rol []model.Role
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	results := r.Db.Find(&rol)
	helper.ErrorPanic(results.Error)
	return rol, nil
}

<<<<<<< HEAD
// FindById implements RolesRepository
func (r *RolesRepositoryImpl) FindById(id int64) (model.Roles, error) {
	var rol model.Roles
	result := r.Db.Find(&rol, id)
	if result != nil {
		return rol, errors.New("roles is not found")
=======
// FindById implements RoleRepository
func (r *RoleRepositoryImpl) FindById(id uint) (model.Role, error) {
	var rol model.Role
	result := r.Db.Find(&rol, id)
	if result != nil {
		return rol, errors.New("Role is not found")
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9

	}
	return rol, nil
}

<<<<<<< HEAD
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

	var updateRole = model.Roles{
		ID:         updatedRoles.ID,
		Name:       updatedRoles.Name,
		Created_at: created_at,
	}
	result := r.Db.Model(&updatedRoles).Updates(updateRole)
	helper.ErrorPanic(result.Error)
	return updatedRoles, nil
}

// FindByName implements RolesRepository
func (r *RolesRepositoryImpl) FindByName(name string) (model.Roles, error) {
	var rol model.Roles
=======
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
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	result := r.Db.First(&rol, "name = ?", name)

	if result.Error != nil {
		return rol, errors.New("invalid name")
	}
	return rol, nil
}
