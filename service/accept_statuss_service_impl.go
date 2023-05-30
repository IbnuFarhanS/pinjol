package service

import (
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/go-playground/validator/v10"
)

type AcceptStatusServiceImpl struct {
	AcceptStatusRepository repository.AcceptStatusRepository
}

// Delete implements BorrowerService
func (s *AcceptStatusServiceImpl) Delete(id int64) (model.AcceptStatus, error) {
	return s.AcceptStatusRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *AcceptStatusServiceImpl) FindAll() ([]model.AcceptStatus, error) {
	return s.AcceptStatusRepository.FindAll()
}

// FindById implements BorrowerService
func (s *AcceptStatusServiceImpl) FindById(id int64) (model.AcceptStatus, error) {
	return s.AcceptStatusRepository.FindById(id)
}

// Save implements BorrowerService
func (s *AcceptStatusServiceImpl) Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {

	newAs := model.AcceptStatus{
		Transactions: newAcceptStatus.Transactions,
		Status:       newAcceptStatus.Status,
		Created_At:   newAcceptStatus.Created_At,
	}
	return s.AcceptStatusRepository.Save(newAs)

}

// Update implements BorrowerService
func (s *AcceptStatusServiceImpl) Update(updateAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {

	var ast model.AcceptStatus
	create_at := ast.Created_At

	newAs := model.AcceptStatus{
		ID:           updateAcceptStatus.ID,
		Transactions: updateAcceptStatus.Transactions,
		Status:       updateAcceptStatus.Status,
		Created_At:   create_at,
	}

	return s.AcceptStatusRepository.Update(newAs)
}

func NewAcceptStatusServiceImpl(acceptStatusRepository repository.AcceptStatusRepository, validate *validator.Validate) AcceptStatusService {
	return &AcceptStatusServiceImpl{
		AcceptStatusRepository: acceptStatusRepository,
	}
}
