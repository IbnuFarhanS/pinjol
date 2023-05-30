package repository

import "github.com/IbnuFarhanS/pinjol/model"

type TransactionsRepository interface {
	Save(newTransactions model.Transactions) (model.Transactions, error)
	Update(updateTransactions model.Transactions) (model.Transactions, error)
	Delete(id int64) (model.Transactions, error)
	FindById(id int64) (model.Transactions, error)
	FindAll() ([]model.Transactions, error)
	FindByUserID(userID int64) ([]model.Transactions, error)
}
