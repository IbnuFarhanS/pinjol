package service

import (
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

type AcceptStatusServiceImpl struct {
	AcceptStatusRepository repository.AcceptStatusRepository
}

// Delete implements BorrowerService
func (s *AcceptStatusServiceImpl) Delete(id uint) (model.AcceptStatus, error) {
	return s.AcceptStatusRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *AcceptStatusServiceImpl) FindAll() ([]model.AcceptStatus, error) {
	return s.AcceptStatusRepository.FindAll()
}

// FindById implements BorrowerService
func (s *AcceptStatusServiceImpl) FindById(id uint) (model.AcceptStatus, error) {
	return s.AcceptStatusRepository.FindById(id)
}

// Save implements BorrowerService
func (s *AcceptStatusServiceImpl) Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {

	newAs := model.AcceptStatus{
		Transaction: model.Transaction{},
		Status:      newAcceptStatus.Status,
		CreatedAt:   time.Time{},
	}
	return s.AcceptStatusRepository.Save(newAs)

}

// Update implements BorrowerService
func (s *AcceptStatusServiceImpl) Update(updateAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {

	var ast model.AcceptStatus
	create_at := ast.CreatedAt

	newAs := model.AcceptStatus{
		ID:            updateAcceptStatus.ID,
		TransactionID: 0,
		Transaction:   model.Transaction{},
		Status:        updateAcceptStatus.Status,
		CreatedAt:     create_at,
	}

	return s.AcceptStatusRepository.Update(newAs)
}

func NewAcceptStatusServiceImpl(acceptStatusRepository repository.AcceptStatusRepository) AcceptStatusService {
	return &AcceptStatusServiceImpl{
		AcceptStatusRepository: acceptStatusRepository,
	}
}
