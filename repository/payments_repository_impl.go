package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

type PaymentsRepositoryImpl struct {
	Db *gorm.DB
}

func NewPaymentsRepositoryImpl(Db *gorm.DB) PaymentsRepository {
	return &PaymentsRepositoryImpl{Db: Db}
}

// Delete implements PaymentsRepository
func (r *PaymentsRepositoryImpl) Delete(id int64) (model.Payments, error) {
	var pay model.Payments
	result := r.Db.Where("id = ?", id).Delete(&pay)
	helper.ErrorPanic(result.Error)
	return pay, nil
}

// FindAll implements PaymentsRepository
func (r *PaymentsRepositoryImpl) FindAll() ([]model.Payments, error) {
	var pay []model.Payments
	results := r.Db.Find(&pay)
	helper.ErrorPanic(results.Error)
	return pay, nil
}

// FindById implements PaymentsRepository
func (r *PaymentsRepositoryImpl) FindById(id int64) (model.Payments, error) {
	var pay model.Payments
	result := r.Db.First(&pay, "id = ?", id)
	if result.Error != nil {
		return pay, errors.New("payments is not found")
	}
	return pay, nil
}

// Save implements PaymentsRepository
func (r *PaymentsRepositoryImpl) Save(newPayments model.Payments) (model.Payments, error) {

	tx := r.Db.Begin()
	var tra model.Transactions
	if err := tx.Table("transactions").
		Where("id = ?", newPayments.TransactionsID).
		First(&tra).Error; err != nil {
		tx.Rollback()
		return model.Payments{}, fmt.Errorf("transactions with id %d not found", newPayments.TransactionsID)
	}

	newPayments.Payment_Date = time.Now()
	newPayments.NextInstallment = tra.Total - newPayments.Payment_Amount

	if err := tx.Table("payments").Create(&newPayments).Error; err != nil {
		tx.Rollback()
		return model.Payments{}, errors.New("failed to save transactions")
	}
	return newPayments, tx.Commit().Error

}

// Update implements PaymentsRepository
func (r *PaymentsRepositoryImpl) Update(updatedPayments model.Payments) (model.Payments, error) {
	result := r.Db.Model(&model.Payments{}).Where("id = ?", updatedPayments.ID).Updates(updatedPayments)
	helper.ErrorPanic(result.Error)
	return updatedPayments, nil
}
