package service

import (
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/go-playground/validator/v10"
)

type TransactionsServiceImpl struct {
	TransactionsRepository repository.TransactionsRepository
	Validate               *validator.Validate
}

// Delete implements BorrowerService
func (s *TransactionsServiceImpl) Delete(id int64) (model.Transactions, error) {
	return s.TransactionsRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *TransactionsServiceImpl) FindAll() ([]model.Transactions, error) {
	return s.TransactionsRepository.FindAll()
}

// FindById implements BorrowerService
func (s *TransactionsServiceImpl) FindById(id int64) (model.Transactions, error) {
	return s.TransactionsRepository.FindById(id)
}

// Save implements BorrowerService
func (s *TransactionsServiceImpl) Save(newTransactions model.Transactions) (model.Transactions, error) {

	newTra := model.Transactions{
		Products:   newTransactions.Products,
		Users:      newTransactions.Users,
		Status:     newTransactions.Status,
		Created_At: newTransactions.Created_At,
		Due_Date:   newTransactions.Due_Date,
	}
	return s.TransactionsRepository.Save(newTra)

}

// Update implements BorrowerService
func (s *TransactionsServiceImpl) Update(updateTransactions model.Transactions) (model.Transactions, error) {

	var tra model.Transactions
	create_at := tra.Created_At

	newTra := model.Transactions{
		ID:         updateTransactions.ID,
		Products:   updateTransactions.Products,
		Users:      updateTransactions.Users,
		Status:     updateTransactions.Status,
		Due_Date:   updateTransactions.Due_Date,
		Created_At: create_at,
	}

	return s.TransactionsRepository.Update(newTra)
}

func NewTransactionsServiceImpl(transactionsRepository repository.TransactionsRepository, validate *validator.Validate) TransactionsService {
	return &TransactionsServiceImpl{
		TransactionsRepository: transactionsRepository,
		Validate:               validate,
	}
}
