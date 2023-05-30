package service

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/go-playground/validator/v10"
)

type TransactionsServiceImpl struct {
	TransactionsRepository repository.TransactionsRepository
	UsersRepository        repository.UsersRepository
}

// Delete implements BorrowerService
func (s *TransactionsServiceImpl) Delete(id int64) (model.Transactions, error) {
	return s.TransactionsRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *TransactionsServiceImpl) FindAll() ([]model.Transactions, error) {
	transactions, err := s.TransactionsRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// FindById implements BorrowerService
func (s *TransactionsServiceImpl) FindById(id int64) (model.Transactions, error) {
	return s.TransactionsRepository.FindById(id)
}

// Save implements BorrowerService
func (s *TransactionsServiceImpl) Save(newTransactions model.Transactions, userid int64) (model.Transactions, error) {

	user, err := s.UsersRepository.FindById(userid)
	if err != nil {
		return model.Transactions{}, err
	}

	if newTransactions.Amount > user.Limit {
		return model.Transactions{}, errors.New("Amount exceeds user's limit")
	}

	user.Limit -= newTransactions.Amount

	_, err = s.UsersRepository.Update(user)
	if err != nil {
		return model.Transactions{}, err
	}

	created_at := time.Now()
	due_date := created_at.AddDate(0, 1, 0)
	newTra := model.Transactions{
		ProductsID: newTransactions.ProductsID,
		UsersID:    userid,
		Users:      user,
		Status:     false,
		Amount:     newTransactions.Amount,
		Created_At: created_at,
		Due_Date:   due_date,
	}
	transaction, err := s.TransactionsRepository.Save(newTra)
	if err != nil {
		return model.Transactions{}, err
	}

	return transaction, nil
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
		Amount:     updateTransactions.Amount,
		Due_Date:   updateTransactions.Due_Date,
		Created_At: create_at,
	}

	return s.TransactionsRepository.Update(newTra)
}

func NewTransactionsServiceImpl(transactionsRepository repository.TransactionsRepository, validate *validator.Validate, usersRepo repository.UsersRepository) TransactionsService {
	return &TransactionsServiceImpl{
		TransactionsRepository: transactionsRepository,
		UsersRepository:        usersRepo,
	}
}
