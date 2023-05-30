package service

import (
<<<<<<< HEAD
	"errors"
=======
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

<<<<<<< HEAD
type RolesServiceImpl struct {
	RolesRepository repository.RolesRepository
}

// Delete implements RolesService
func (s *RolesServiceImpl) Delete(id int64) (model.Roles, error) {
	return s.RolesRepository.Delete(id)
}

// FindAll implements RolesService
func (s *RolesServiceImpl) FindAll() ([]model.Roles, error) {
	return s.RolesRepository.FindAll()
}

// FindById implements RolesService
func (s *RolesServiceImpl) FindById(id int64) (model.Roles, error) {
	return s.RolesRepository.FindById(id)
}

// FindByUsername implements RolesService
func (s *RolesServiceImpl) FindByName(name string) (model.Roles, error) {
	return s.RolesRepository.FindByName(name)
}

// Save implements RolesService
func (s *RolesServiceImpl) Save(newRoles model.Roles) (model.Roles, error) {
	// Validate name
	if newRoles.Name == "" {
		return model.Roles{}, errors.New("name is required")
	}

	created_at := time.Now()
	newBor := model.Roles{
		Name:       newRoles.Name,
		Created_at: created_at,
	}
	return s.RolesRepository.Save(newBor)
}

// Update implements RolesService
func (s *RolesServiceImpl) Update(updatedRoles model.Roles) (model.Roles, error) {
	// Validate name
	if updatedRoles.Name == "" {
		return model.Roles{}, errors.New("name is required")
	}

	var pay model.Roles
	create_at := pay.Created_at

	newRol := model.Roles{
		ID:         updatedRoles.ID,
		Name:       updatedRoles.Name,
		Created_at: create_at,
	}

	return s.RolesRepository.Update(newRol)
}

func NewRolesServiceImpl(RolesRepository repository.RolesRepository) RolesService {
	return &RolesServiceImpl{
		RolesRepository: RolesRepository,
=======
type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
}

// Delete implements RoleService
func (s *RoleServiceImpl) Delete(id uint) (model.Role, error) {
	return s.RoleRepository.Delete(id)
}

// FindAll implements RoleService
func (s *RoleServiceImpl) FindAll() ([]model.Role, error) {
	return s.RoleRepository.FindAll()
}

// FindById implements RoleService
func (s *RoleServiceImpl) FindById(id uint) (model.Role, error) {
	return s.RoleRepository.FindById(id)
}

// FindByUsername implements RoleService
func (s *RoleServiceImpl) FindByName(name string) (model.Role, error) {
	return s.RoleRepository.FindByName(name)
}

// Save implements RoleService
func (s *RoleServiceImpl) Save(newRole model.Role) (model.Role, error) {
	created_at := time.Now()
	newBor := model.Role{
		Name:      newRole.Name,
		CreatedAt: created_at,
	}
	return s.RoleRepository.Save(newBor)
}

// Update implements RoleService
func (s *RoleServiceImpl) Update(updatedRole model.Role) (model.Role, error) {
	panic("unimplemented")
}

func NewRoleServiceImpl(RoleRepository repository.RoleRepository) RoleService {
	return &RoleServiceImpl{
		RoleRepository: RoleRepository,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}
}
