package service

import (
	"time"

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
func (s *TransactionsServiceImpl) Save(newTransactions model.Transactions, userid int64) (model.Transactions, error) {

	created_at := time.Now()
	newTra := model.Transactions{
		ProductsID: newTransactions.ProductsID,
		UsersID:    userid,
		Status:     false,
		Created_At: created_at,
		Due_Date:   created_at,
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

func NewTransactionsServiceImpl(TransactionsRepository repository.TransactionsRepository, validate *validator.Validate) TransactionsService {
	return &TransactionsServiceImpl{
		TransactionsRepository: TransactionsRepository,
		Validate:               validate,
	}
}
