package service

import (
<<<<<<< HEAD
	"errors"
=======
	"time"
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

type AcceptStatusServiceImpl struct {
	AcceptStatusRepository repository.AcceptStatusRepository
<<<<<<< HEAD
}

// Delete implements AcceptStatus
func (s *AcceptStatusServiceImpl) Delete(id int64) (model.AcceptStatus, error) {
	return s.AcceptStatusRepository.Delete(id)
}

// FindAll implements AcceptStatus
=======
	TransactionRepository  repository.TransactionRepository
}

// Delete implements BorrowerService
func (s *AcceptStatusServiceImpl) Delete(id uint) (model.AcceptStatus, error) {
	return s.AcceptStatusRepository.Delete(id)
}

// FindAll implements BorrowerService
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
func (s *AcceptStatusServiceImpl) FindAll() ([]model.AcceptStatus, error) {
	return s.AcceptStatusRepository.FindAll()
}

<<<<<<< HEAD
// FindById implements AcceptStatus
func (s *AcceptStatusServiceImpl) FindById(id int64) (model.AcceptStatus, error) {
	return s.AcceptStatusRepository.FindById(id)
}

// Save implements AcceptStatus
func (s *AcceptStatusServiceImpl) Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	// Validate id_transaction
	if newAcceptStatus.TransactionsID == 0 {
		return model.AcceptStatus{}, errors.New("id_transaction is required")
	}

	newAs := model.AcceptStatus{
		TransactionsID: newAcceptStatus.TransactionsID,
		Status:         newAcceptStatus.Status,
		Created_At:     newAcceptStatus.Created_At,
	}
	return s.AcceptStatusRepository.Save(newAs)

}

// Update implements AcceptStatus
func (s *AcceptStatusServiceImpl) Update(updateAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	// Validate id_transaction
	if updateAcceptStatus.TransactionsID == 0 {
		return model.AcceptStatus{}, errors.New("id_transaction tidak boleh kosong")
	}

	var ast model.AcceptStatus
	create_at := ast.Created_At

	newAs := model.AcceptStatus{
		ID:             updateAcceptStatus.ID,
		TransactionsID: updateAcceptStatus.TransactionsID,
		Status:         updateAcceptStatus.Status,
		Created_At:     create_at,
=======
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
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}

	return s.AcceptStatusRepository.Update(newAs)
}

<<<<<<< HEAD
func NewAcceptStatusServiceImpl(AcceptStatusRepository repository.AcceptStatusRepository) AcceptStatusService {
	return &AcceptStatusServiceImpl{
		AcceptStatusRepository: AcceptStatusRepository,
=======
func NewAcceptStatusServiceImpl(acceptStatusRepository repository.AcceptStatusRepository, TransactionRepository repository.TransactionRepository) AcceptStatusService {
	return &AcceptStatusServiceImpl{
		AcceptStatusRepository: acceptStatusRepository,
		TransactionRepository:  TransactionRepository,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}
}
