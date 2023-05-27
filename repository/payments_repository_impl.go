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
	var bor model.Payments
	result := r.Db.Where("id = ?", id).Delete(&bor)
	helper.ErrorPanic(result.Error)
	return bor, nil
}

// FindAll implements PaymentsRepository
func (r *PaymentsRepositoryImpl) FindAll() ([]model.Payments, error) {
	var bor []model.Payments
	results := r.Db.Find(&bor)
	helper.ErrorPanic(results.Error)
	return bor, nil
}

// FindById implements PaymentsRepository
func (r *PaymentsRepositoryImpl) FindById(id int64) (model.Payments, error) {
	var bor model.Payments
	result := r.Db.Find(&bor, "id = ?", id)
	if result.Error != nil {
		return bor, errors.New("payments is not found")
	}
	return bor, nil
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
	result := r.Db.Model(&model.Payments{}).Where("id = ?", updatedPayments.ID).Updates(updatedPayments)
	helper.ErrorPanic(result.Error)
	return updatedPayments, nil
}
