package service

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

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
	}
}
