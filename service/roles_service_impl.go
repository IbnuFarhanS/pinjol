package service

import (
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	reposity "github.com/IbnuFarhanS/pinjol/repository"
	"github.com/go-playground/validator/v10"
)

type RolesServiceImpl struct {
	RolesRepository reposity.RolesRepository
	Validate        *validator.Validate
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
func (s *RolesServiceImpl) FindByUsername(username string) (model.Roles, error) {
	return s.RolesRepository.FindByUsername(username)
}

// Save implements RolesService
func (s *RolesServiceImpl) Save(newRoles model.Roles) (model.Roles, error) {
	created_at := time.Now()
	newBor := model.Roles{
		Name:       newRoles.Name,
		Created_at: created_at,
	}
	return s.RolesRepository.Save(newBor)
}

// Update implements RolesService
func (s *RolesServiceImpl) Update(updatedRoles model.Roles) (model.Roles, error) {
	panic("unimplemented")
}

func NewRolesServiceImpl(RolesRepository reposity.RolesRepository, validate *validator.Validate) RolesService {
	return &RolesServiceImpl{
		RolesRepository: RolesRepository,
		Validate:        validate,
	}
}
