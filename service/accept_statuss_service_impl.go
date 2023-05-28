package service

import (
	"errors"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/go-playground/validator/v10"
)

type AcceptStatusServiceImpl struct {
	AcceptStatusRepository repository.AcceptStatusRepository
	Validate               *validator.Validate
}

// Delete implements AcceptStatus
func (s *AcceptStatusServiceImpl) Delete(id int64) (model.AcceptStatus, error) {
	return s.AcceptStatusRepository.Delete(id)
}

// FindAll implements AcceptStatus
func (s *AcceptStatusServiceImpl) FindAll() ([]model.AcceptStatus, error) {
	return s.AcceptStatusRepository.FindAll()
}

// FindById implements AcceptStatus
func (s *AcceptStatusServiceImpl) FindById(id int64) (model.AcceptStatus, error) {
	return s.AcceptStatusRepository.FindById(id)
}

// Save implements AcceptStatus
func (s *AcceptStatusServiceImpl) Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	if newAcceptStatus.TransactionsID == 0 {
		return model.AcceptStatus{}, errors.New("id_transaction tidak boleh kosong")
	}

	newAs := model.AcceptStatus{
		Transactions: newAcceptStatus.Transactions,
		Status:       newAcceptStatus.Status,
		Created_At:   newAcceptStatus.Created_At,
	}
	return s.AcceptStatusRepository.Save(newAs)

}

// Update implements AcceptStatus
func (s *AcceptStatusServiceImpl) Update(updateAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	if updateAcceptStatus.TransactionsID == 0 {
		return model.AcceptStatus{}, errors.New("id_transaction tidak boleh kosong")
	}

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

func NewAcceptStatusServiceImpl(AcceptStatusRepository repository.AcceptStatusRepository, validate *validator.Validate) AcceptStatusService {
	return &AcceptStatusServiceImpl{
		AcceptStatusRepository: AcceptStatusRepository,
		Validate:               validate,
	}
}
