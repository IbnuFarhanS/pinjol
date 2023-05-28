package service

import "github.com/IbnuFarhanS/pinjol/model"

type TransactionsService interface {
	Save(newTransactions model.Transactions, userid int64) (model.Transactions, error)
	Update(updateTransactions model.Transactions) (model.Transactions, error)
	Delete(id int64) (model.Transactions, error)
	FindById(id int64) (model.Transactions, error)
	FindAll() ([]model.Transactions, error)
}
