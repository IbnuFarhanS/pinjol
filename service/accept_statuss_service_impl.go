package service

import (
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

type AcceptStatusServiceImpl struct {
	AcceptStatusRepository repository.AcceptStatusRepository
	TransactionRepository  repository.TransactionRepository
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
	transaction := newAcceptStatus.Transaction

	if newAcceptStatus.Status {
		transaction.Status = true
	} else {
		transaction.Status = false
	}

	currentTime := time.Now()
	newAcceptStatus.CreatedAt = currentTime
	newAcceptStatus.Transaction = transaction

	// Update status Transaction berdasarkan ID transaksi
	err := s.TransactionRepository.UpdateStatus(transaction.ID, transaction.Status)
	if err != nil {
		return model.AcceptStatus{}, err
	}

	return s.AcceptStatusRepository.Save(newAcceptStatus)
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

func NewAcceptStatusServiceImpl(acceptStatusRepository repository.AcceptStatusRepository, TransactionRepository repository.TransactionRepository) AcceptStatusService {
	return &AcceptStatusServiceImpl{
		AcceptStatusRepository: acceptStatusRepository,
		TransactionRepository:  TransactionRepository,
	}
}
