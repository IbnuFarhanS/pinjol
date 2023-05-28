package repository

import (
	"errors"
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
	result := r.Db.Find(&pay, "id = ?", id)
	if result.Error != nil {
		return pay, errors.New("payments is not found")
	}
	return pay, nil
}

// Save implements PaymentsRepository
func (r *PaymentsRepositoryImpl) Save(newPayments model.Payments) (model.Payments, error) {
	currentTime := time.Now()
	newPayments.Payment_Date = currentTime
	result := r.Db.Create(&newPayments)
	helper.ErrorPanic(result.Error)
	return newPayments, nil
}

// Update implements PaymentsRepository
func (r *PaymentsRepositoryImpl) Update(updatedPayments model.Payments) (model.Payments, error) {
	var rol model.Payments
	created_at := rol.Payment_Date

	var updatePayment = model.Payments{
		ID:              updatedPayments.ID,
		TransactionsID:  updatedPayments.TransactionsID,
		PaymentMethodID: updatedPayments.PaymentMethodID,
		Payment_Amount:  updatedPayments.Payment_Amount,
		Payment_Date:    created_at,
	}
	result := r.Db.Model(&updatedPayments).Updates(updatePayment)
	helper.ErrorPanic(result.Error)
	return updatedPayments, nil
}
