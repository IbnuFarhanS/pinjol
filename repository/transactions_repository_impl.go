package repository

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	Db *gorm.DB
}

func NewTransactionRepositoryImpl(Db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{Db: Db}
}

// Delete implements TransactionRepository
func (r *TransactionRepositoryImpl) Delete(id uint) (model.Transaction, error) {
	var tra model.Transaction
	result := r.Db.Where("id = ?", id).Delete(&tra)
	helper.ErrorPanic(result.Error)
	return tra, nil
}

// FindAll implements TransactionRepository
func (r *TransactionRepositoryImpl) FindAll() ([]model.Transaction, error) {
	var tra []model.Transaction

	result := r.Db.Model(&tra).Select("Transaction.id_user, Transaction.id_product, Transaction.status, Transaction.created_at, Transaction.due_date, Transaction.amount, users.id, products.id").Joins("join products on products.id = Transaction.id_product").Joins("join users on users.id = Transaction.id_user").Scan(&tra)
	if result.Error != nil {
		return nil, result.Error
	}
	return tra, nil
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
	currentTime := time.Now()
	newTransaction.CreatedAt = currentTime
	result := r.Db.Create(&newTransaction)
	helper.ErrorPanic(result.Error)
	return newTransaction, nil

}

// Update implements TransactionRepository
func (r *TransactionRepositoryImpl) Update(updatedTransaction model.Transaction) (model.Transaction, error) {
	result := r.Db.Model(&model.Transaction{}).Where("id = ?", updatedTransaction.ID).Updates(updatedTransaction)
	helper.ErrorPanic(result.Error)
	return updatedTransaction, nil
}

func (r *TransactionRepositoryImpl) FindByUserID(userID uint) ([]model.Transaction, error) {
	var Transaction []model.Transaction
	if err := r.Db.Where("user_id = ?", userID).Find(&Transaction).Error; err != nil {
		return nil, err
	}
	return Transaction, nil
}
