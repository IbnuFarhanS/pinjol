package service

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
<<<<<<< HEAD
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
	// Validate id_product
	if newTransactions.ProductsID == 0 {
		return model.Transactions{}, errors.New("id product is required")
	}
	// Validate id_user
	if newTransactions.UsersID == 0 {
		return model.Transactions{}, errors.New("id user is required")
	}

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
	// Validate id_product
	if updateTransactions.ProductsID == 0 {
		return model.Transactions{}, errors.New("id product is required")
	}
	// Validate id_user
	if updateTransactions.UsersID == 0 {
		return model.Transactions{}, errors.New("id user is required")
	}

	var tra model.Transactions
	create_at := tra.Created_At

	newTra := model.Transactions{
		ID:         updateTransactions.ID,
		ProductsID: updateTransactions.ProductsID,
		UsersID:    updateTransactions.UsersID,
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
=======
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	UserRepository        repository.UserRepository
	// ProductRepository     repository.ProductRepository
	// ProductService        ProductService
}

// Delete implements BorrowerService
func (s *TransactionServiceImpl) Delete(id uint) (model.Transaction, error) {
	return s.TransactionRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *TransactionServiceImpl) FindAll() ([]model.Transaction, error) {
	Transaction, err := s.TransactionRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return Transaction, nil
}

// FindById implements BorrowerService
func (s *TransactionServiceImpl) FindById(id uint) (model.Transaction, error) {
	return s.TransactionRepository.FindById(id)
}

// Save implements BorrowerService
func (s *TransactionServiceImpl) Save(newTransaction model.Transaction, userid uint) (model.Transaction, error) {

	user, err := s.UserRepository.FindById(userid)
	if err != nil {
		return model.Transaction{}, err
	}

	if newTransaction.Amount > user.Limit {
		return model.Transaction{}, errors.New("Amount exceeds user's limit")
	}

	user.Limit -= newTransaction.Amount

	_, err = s.UserRepository.Update(user)
	if err != nil {
		return model.Transaction{}, err
	}

	created_at := time.Now()
	due_date := created_at.AddDate(0, 1, 0)
	newTra := model.Transaction{
		ProductID: newTransaction.ProductID,
		UserID:    userid,
		User:      user,
		Status:    false,
		Amount:    newTransaction.Amount,
		CreatedAt: created_at,
		DueDate:   due_date,
	}
	transaction, err := s.TransactionRepository.Save(newTra)
	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

// Update implements BorrowerService
func (s *TransactionServiceImpl) Update(updateTransaction model.Transaction) (model.Transaction, error) {

	var tra model.Transaction
	create_at := tra.CreatedAt
	due_date := tra.DueDate

	newTra := model.Transaction{
		ID:        updateTransaction.ID,
		Product:   updateTransaction.Product,
		User:      updateTransaction.User,
		Status:    updateTransaction.Status,
		Amount:    updateTransaction.Amount,
		DueDate:   due_date,
		CreatedAt: create_at,
	}

	return s.TransactionRepository.Update(newTra)
}

func NewTransactionServiceImpl(TransactionRepository repository.TransactionRepository, UserRepo repository.UserRepository) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: TransactionRepository,
		UserRepository:        UserRepo,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}
}
