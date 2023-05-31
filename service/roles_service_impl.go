package service

import (
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

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
	var rol model.Role
	create_at := rol.CreatedAt

	newRol := model.Role{
		ID:        updatedRole.ID,
		Name:      updatedRole.Name,
		CreatedAt: create_at,
	}

	return s.RoleRepository.Update(newRol)
}

func NewRoleServiceImpl(RoleRepository repository.RoleRepository) RoleService {
	return &RoleServiceImpl{
		RoleRepository: RoleRepository,
	}
}
