package repository

import (
	"errors"
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
		return pay, errors.New("Payment is not found")
	}
	return pay, nil
}

// Save implements PaymentRepository
func (r *PaymentRepositoryImpl) Save(newPayment model.Payment) (model.Payment, error) {
	currentTime := time.Now()
	newPayment.PaymentDate = currentTime
	result := r.Db.Create(&newPayment)
	helper.ErrorPanic(result.Error)
	return newPayment, nil
}

// Update implements PaymentRepository
func (r *PaymentRepositoryImpl) Update(updatedPayment model.Payment) (model.Payment, error) {
	result := r.Db.Model(&model.Payment{}).Where("id = ?", updatedPayment.ID).Updates(updatedPayment)
	helper.ErrorPanic(result.Error)
	return updatedPayment, nil
}
