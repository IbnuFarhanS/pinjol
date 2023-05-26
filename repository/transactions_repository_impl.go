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
	var bor model.Transactions
	result := r.Db.Where("id = ?", id).Delete(&bor)
	helper.ErrorPanic(result.Error)
	return bor, nil
}

// FindAll implements TransactionsRepository
func (r *TransactionsRepositoryImpl) FindAll() ([]model.Transactions, error) {
	var bor []model.Transactions
	results := r.Db.Find(&bor)
	helper.ErrorPanic(results.Error)
	return bor, nil
}

// FindById implements TransactionsRepository
func (r *TransactionsRepositoryImpl) FindById(id int64) (model.Transactions, error) {
	var bor model.Transactions
	result := r.Db.Find(&bor, "id = ?", id)
	if result.Error != nil {
		return bor, errors.New("transactions is not found")
	}
	return bor, nil
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
