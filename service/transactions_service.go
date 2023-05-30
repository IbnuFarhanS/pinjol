package service

import "github.com/IbnuFarhanS/pinjol/model"

type TransactionService interface {
	Save(newTransaction model.Transaction, userid uint) (model.Transaction, error)
	Update(updateTransaction model.Transaction) (model.Transaction, error)
	Delete(id uint) (model.Transaction, error)
	FindById(id uint) (model.Transaction, error)
	FindAll() ([]model.Transaction, error)
}
