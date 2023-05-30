package repository

import (
	"errors"
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

<<<<<<< HEAD
type PaymentsRepositoryImpl struct {
	Db *gorm.DB
}

func NewPaymentsRepositoryImpl(Db *gorm.DB) PaymentsRepository {
	return &PaymentsRepositoryImpl{Db: Db}
}

// Delete implements PaymentsRepository
func (r *PaymentsRepositoryImpl) Delete(id int64) (model.Payments, error) {
	var pay model.Payments
=======
type PaymentRepositoryImpl struct {
	Db *gorm.DB
}

func NewPaymentRepositoryImpl(Db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{Db: Db}
}

// Delete implements PaymentRepository
func (r *PaymentRepositoryImpl) Delete(id uint) (model.Payment, error) {
	var pay model.Payment
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	result := r.Db.Where("id = ?", id).Delete(&pay)
	helper.ErrorPanic(result.Error)
	return pay, nil
}

<<<<<<< HEAD
// FindAll implements PaymentsRepository
func (r *PaymentsRepositoryImpl) FindAll() ([]model.Payments, error) {
	var pay []model.Payments
=======
// FindAll implements PaymentRepository
func (r *PaymentRepositoryImpl) FindAll() ([]model.Payment, error) {
	var pay []model.Payment
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	results := r.Db.Find(&pay)
	helper.ErrorPanic(results.Error)
	return pay, nil
}

<<<<<<< HEAD
// FindById implements PaymentsRepository
func (r *PaymentsRepositoryImpl) FindById(id int64) (model.Payments, error) {
	var pay model.Payments
	result := r.Db.Find(&pay, "id = ?", id)
	if result.Error != nil {
		return pay, errors.New("payments is not found")
=======
// FindById implements PaymentRepository
func (r *PaymentRepositoryImpl) FindById(id uint) (model.Payment, error) {
	var pay model.Payment
	result := r.Db.First(&pay, "id = ?", id)
	if result.Error != nil {
		return pay, errors.New("Payment is not found")
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}
	return pay, nil
}

<<<<<<< HEAD
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
=======
// Save implements PaymentRepository
func (r *PaymentRepositoryImpl) Save(newPayment model.Payment) (model.Payment, error) {
	tx := r.Db.Begin()

	var tra model.Transaction
	if err := tx.Table("transactions").
		Where("id = ?", newPayment.TransactionID).
		First(&tra).Error; err != nil {
		tx.Rollback()
		return model.Payment{}, fmt.Errorf("transactions with id %d not found", newPayment.TransactionID)
	}

	newPayment.PaymentDate = time.Now()
	newPayment.NextInstallment = tra.Total - newPayment.PaymentAmount

	if err := tx.Table("payments").Create(&newPayment).Error; err != nil {
		tx.Rollback()
		return model.Payment{}, fmt.Errorf("failed to save payments: %v", err)
	}

	return newPayment, tx.Commit().Error
}

// Update implements PaymentRepository
func (r *PaymentRepositoryImpl) Update(updatedPayment model.Payment) (model.Payment, error) {
	result := r.Db.Model(&model.Payment{}).Where("id = ?", updatedPayment.ID).Updates(updatedPayment)
	helper.ErrorPanic(result.Error)
	return updatedPayment, nil
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
