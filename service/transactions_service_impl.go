package service

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
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
	if newTransaction.Amount == 0 {
		return model.Transaction{}, errors.New("amoun is required")
	}
	user, err := s.UserRepository.FindById(userid)
	if err != nil {
		return model.Transaction{}, err
	}

	if newTransaction.Amount > user.Limit {
		return model.Transaction{}, errors.New("amount exceeds user's limit")
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
	}
}
