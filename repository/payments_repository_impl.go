package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
	Db *gorm.DB
}

func NewPaymentRepositoryImpl(Db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{Db: Db}
}

// Delete implements PaymentRepository
func (r *PaymentRepositoryImpl) Delete(id uint) (model.Payment, error) {
	var pay model.Payment
	result := r.Db.Where("id = ?", id).Delete(&pay)
	helper.ErrorPanic(result.Error)
	return pay, nil
}

// FindAll implements PaymentRepository
func (r *PaymentRepositoryImpl) FindAll() ([]model.Payment, error) {
	var pay []model.Payment
	results := r.Db.Find(&pay)
	helper.ErrorPanic(results.Error)
	return pay, nil
}

// FindById implements PaymentRepository
func (r *PaymentRepositoryImpl) FindById(id uint) (model.Payment, error) {
	var pay model.Payment
	result := r.Db.First(&pay, "id = ?", id)
	if result.Error != nil {
		return pay, errors.New("payment is not found")
	}
	return pay, nil
}

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
}
