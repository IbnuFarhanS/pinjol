package repository

import (
	"errors"
<<<<<<< HEAD
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
=======
	"fmt"
	"time"

>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

<<<<<<< HEAD
type TransactionsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTransactionsRepositoryImpl(Db *gorm.DB) TransactionsRepository {
	return &TransactionsRepositoryImpl{Db: Db}
}

// Delete implements TransactionsRepository
func (r *TransactionsRepositoryImpl) Delete(id int64) (model.Transactions, error) {
	var tra model.Transactions
	result := r.Db.Where("id = ?", id).Delete(&tra)
	helper.ErrorPanic(result.Error)
	return tra, nil
}

// FindAll implements TransactionsRepository
func (r *TransactionsRepositoryImpl) FindAll() ([]model.Transactions, error) {
	var tra []model.Transactions
	results := r.Db.Find(&tra)
	helper.ErrorPanic(results.Error)
	return tra, nil
}

// FindById implements TransactionsRepository
func (r *TransactionsRepositoryImpl) FindById(id int64) (model.Transactions, error) {
	var tra model.Transactions
	result := r.Db.Find(&tra, "id = ?", id)
	if result.Error != nil {
		return tra, errors.New("transactions is not found")
=======
type TransactionRepositoryImpl struct {
	Db *gorm.DB
}

// UpdateStatus implements TransactionRepository
func (r *TransactionRepositoryImpl) UpdateStatus(transactionID uint, status bool) error {
	result := r.Db.Model(&model.Transaction{}).Where("id = ?", transactionID).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewTransactionRepositoryImpl(Db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{Db: Db}
}

// Delete implements TransactionRepository
func (r *TransactionRepositoryImpl) Delete(id uint) (model.Transaction, error) {
	var tra model.Transaction
	result := r.Db.Where("id = ?", id).Delete(&tra)
	if result.Error != nil {
		return tra, errors.New("Transaction is not found")
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}
	return tra, nil
}

<<<<<<< HEAD
// Save implements TransactionsRepository
func (r *TransactionsRepositoryImpl) Save(newTransactions model.Transactions) (model.Transactions, error) {
	currentTime := time.Now()
	newTransactions.Created_At = currentTime
	result := r.Db.Create(&newTransactions)
	helper.ErrorPanic(result.Error)
	return newTransactions, nil
}

// Update implements TransactionsRepository
func (r *TransactionsRepositoryImpl) Update(updatedTransactions model.Transactions) (model.Transactions, error) {
	var rol model.Transactions
	created_at := rol.Created_At

	var updateTra = model.Transactions{
		ID:         updatedTransactions.ID,
		UsersID:    updatedTransactions.UsersID,
		ProductsID: updatedTransactions.ProductsID,
		Status:     updatedTransactions.Status,
		Created_At: created_at,
		Due_Date:   created_at,
	}
	result := r.Db.Model(&updatedTransactions).Updates(updateTra)
	helper.ErrorPanic(result.Error)
	return updatedTransactions, nil
=======
// FindAll implements TransactionRepository
func (r *TransactionRepositoryImpl) FindAll() ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := r.Db.Table("transactions").Find(&transactions).Error; err != nil {
		return nil, err
	}

	for i := range transactions {
		product := model.Product{}
		if err := r.Db.Table("products").Where("id = ?", transactions[i].ProductID).First(&product).Error; err != nil {
			return nil, err
		}

		transactions[i].TotalTax = product.Interest * transactions[i].Amount / 100
		transactions[i].Total = transactions[i].Amount + transactions[i].TotalTax
		transactions[i].TotalMonth = transactions[i].Total / float64(product.Installment)
	}

	return transactions, nil
}

// FindById implements TransactionRepository
func (r *TransactionRepositoryImpl) FindById(id uint) (model.Transaction, error) {
	var tra model.Transaction
	result := r.Db.First(&tra, "id = ?", id)
	if result.Error != nil {
		return tra, errors.New("Transaction is not found")
	}
	return tra, nil
}

// Save implements TransactionRepository
func (r *TransactionRepositoryImpl) Save(newTransaction model.Transaction) (model.Transaction, error) {
	tx := r.Db.Begin()
	var product model.Product
	if err := tx.Table("products").
		Where("id = ?", newTransaction.ProductID).
		First(&product).Error; err != nil {
		tx.Rollback()
		return model.Transaction{}, fmt.Errorf("product with id %d not found", newTransaction.ProductID)
	}

	newTransaction.CreatedAt = time.Now()
	newTransaction.TotalTax = product.Interest * newTransaction.Amount / 100
	newTransaction.Total = newTransaction.Amount + newTransaction.TotalTax
	newTransaction.TotalMonth = newTransaction.Total / float64(product.Installment)

	if err := tx.Table("transactions").Create(&newTransaction).Error; err != nil {
		tx.Rollback()
		return model.Transaction{}, err
	}

	return newTransaction, tx.Commit().Error

}

// Update implements TransactionRepository
func (r *TransactionRepositoryImpl) Update(updatedTransaction model.Transaction) (model.Transaction, error) {
	result := r.Db.Model(&model.Transaction{}).Where("id = ?", updatedTransaction.ID).Updates(updatedTransaction)
	if result.Error != nil {
		return updatedTransaction, errors.New("Transaction is not found")
	}
	return updatedTransaction, nil
}

func (r *TransactionRepositoryImpl) FindByUserID(userID uint) ([]model.Transaction, error) {
	var Transaction []model.Transaction
	if err := r.Db.Where("user_id = ?", userID).Find(&Transaction).Error; err != nil {
		return Transaction, err
	}
	return Transaction, nil
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
