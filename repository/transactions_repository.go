package repository

import "github.com/IbnuFarhanS/pinjol/model"

<<<<<<< HEAD
type TransactionsRepository interface {
	Save(newTransactions model.Transactions) (model.Transactions, error)
	Update(updateTransactions model.Transactions) (model.Transactions, error)
	Delete(id int64) (model.Transactions, error)
	FindById(id int64) (model.Transactions, error)
	FindAll() ([]model.Transactions, error)
=======
type TransactionRepository interface {
	Save(newTransaction model.Transaction) (model.Transaction, error)
	Update(updateTransaction model.Transaction) (model.Transaction, error)
	Delete(id uint) (model.Transaction, error)
	FindById(id uint) (model.Transaction, error)
	FindAll() ([]model.Transaction, error)
	FindByUserID(userID uint) ([]model.Transaction, error)
	UpdateStatus(transactionID uint, status bool) error
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
