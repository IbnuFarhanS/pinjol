package repository

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

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

	result := r.Db.Model(&tra).Select("transactions.id_user, transactions.id_product, transactions.status, transactions.created_at, transactions.due_date, transactions.amount, users.id, products.id").Joins("join products on products.id = transactions.id_product").Joins("join users on users.id = transactions.id_user").Scan(&tra)
	if result.Error != nil {
		return nil, result.Error
	}
	return tra, nil
}

// FindById implements TransactionsRepository
func (r *TransactionsRepositoryImpl) FindById(id int64) (model.Transactions, error) {
	var tra model.Transactions
	result := r.Db.First(&tra, "id = ?", id)
	if result.Error != nil {
		return tra, errors.New("transactions is not found")
	}
	return tra, nil
}

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
	result := r.Db.Model(&model.Transactions{}).Where("id = ?", updatedTransactions.ID).Updates(updatedTransactions)
	helper.ErrorPanic(result.Error)
	return updatedTransactions, nil
}

func (r *TransactionsRepositoryImpl) FindByUserID(userID int64) ([]model.Transactions, error) {
	var transactions []model.Transactions
	if err := r.Db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
