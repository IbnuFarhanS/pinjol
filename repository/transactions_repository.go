package repository

import "github.com/IbnuFarhanS/pinjol/model"

type TransactionRepository interface {
	Save(newTransaction model.Transaction) (model.Transaction, error)
	Update(updateTransaction model.Transaction) (model.Transaction, error)
	Delete(id uint) (model.Transaction, error)
	FindById(id uint) (model.Transaction, error)
	FindAll() ([]model.Transaction, error)
	FindByUserID(userID uint) ([]model.Transaction, error)
}
