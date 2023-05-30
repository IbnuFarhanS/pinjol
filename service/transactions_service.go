package service

import "github.com/IbnuFarhanS/pinjol/model"

<<<<<<< HEAD
type TransactionsService interface {
	Save(newTransactions model.Transactions, userid int64) (model.Transactions, error)
	Update(updateTransactions model.Transactions) (model.Transactions, error)
	Delete(id int64) (model.Transactions, error)
	FindById(id int64) (model.Transactions, error)
	FindAll() ([]model.Transactions, error)
=======
type TransactionService interface {
	Save(newTransaction model.Transaction, userid uint) (model.Transaction, error)
	Update(updateTransaction model.Transaction) (model.Transaction, error)
	Delete(id uint) (model.Transaction, error)
	FindById(id uint) (model.Transaction, error)
	FindAll() ([]model.Transaction, error)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
