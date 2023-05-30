package repository

import (
	"errors"
	"fmt"
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

	var installment []int64
	err := r.Db.Raw(`SELECT p.installment
	FROM transactions tr
	JOIN products p on p.id = tr.id_product
	JOIN users u on u.id = tr.id_user;
	`).Scan(&installment).Error

	if err != nil {
		return tra, err
	}

	err = r.Db.Raw(`
	SELECT
		tr.id as id,
		tr.id_product as id_product, 
		tr.id_user as id_user,
		tr.status as status, 
		tr.amount as amount, 
		tr.created_at as created_at, 
		tr.due_date as due_date, 
		((tr.amount + ((p.bunga * tr.amount)/100)) / p.installment) as total_mounth,
		((p.bunga * tr.amount)/100) as total_tax,
		(tr.amount + ((p.bunga * tr.amount)/100)) as total
	FROM transactions tr
	JOIN products p on p.id = tr.id_product 
	JOIN users u on u.id = tr.id_user;
	`).Scan(&tra).Error
	if err != nil {
		return tra, err
	}
	for i := range tra {
		tra[i].TotalMounth = tra[i].Total / float64(installment[i])
	}

	return tra, nil
}

// FindById implements TransactionsRepository
func (r *TransactionsRepositoryImpl) FindById(id int64) (model.Transactions, error) {
	var tra model.Transactions
	result := r.Db.First(&tra, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return tra, errors.New("transaction is not found")
		}
		return tra, result.Error
	}
	return tra, nil
}

// Save implements TransactionsRepository
func (r *TransactionsRepositoryImpl) Save(newTransactions model.Transactions) (model.Transactions, error) {

	tx := r.Db.Begin()
	var product model.Products
	if err := tx.Table("products").
		Where("id = ?", newTransactions.ProductsID).
		First(&product).Error; err != nil {
		tx.Rollback()
		return model.Transactions{}, fmt.Errorf("product with id %d not found", newTransactions.ProductsID)
	}

	newTransactions.Created_At = time.Now()
	newTransactions.TotalTax = product.Bunga * newTransactions.Amount / 100
	newTransactions.Total = newTransactions.Amount + newTransactions.TotalTax
	newTransactions.TotalMounth = newTransactions.Total / float64(product.Installment)

	if err := tx.Table("transactions").Create(&newTransactions).Error; err != nil {
		tx.Rollback()
		return model.Transactions{}, err
	}

	return newTransactions, tx.Commit().Error
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
