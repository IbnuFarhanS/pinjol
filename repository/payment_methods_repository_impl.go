package repository

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

type PaymentMethodRepositoryImpl struct {
	Db *gorm.DB
}

func NewPaymentMethodRepositoryImpl(Db *gorm.DB) PaymentMethodRepository {
	return &PaymentMethodRepositoryImpl{Db: Db}
}

// Delete implements PaymentMethodRepository
func (r *PaymentMethodRepositoryImpl) Delete(id int64) (model.PaymentMethod, error) {
	var bor model.PaymentMethod
	result := r.Db.Where("id = ?", id).Delete(&bor)
	helper.ErrorPanic(result.Error)
	return bor, nil
}

// FindAll implements PaymentMethodRepository
func (r *PaymentMethodRepositoryImpl) FindAll() ([]model.PaymentMethod, error) {
	var bor []model.PaymentMethod
	results := r.Db.Find(&bor)
	helper.ErrorPanic(results.Error)
	return bor, nil
}

// FindById implements PaymentMethodRepository
func (r *PaymentMethodRepositoryImpl) FindById(id int64) (model.PaymentMethod, error) {
	var bor model.PaymentMethod
	result := r.Db.Find(&bor, "id = ?", id)
	if result.Error != nil {
		return bor, errors.New("paymentMethod is not found")
	}
	return bor, nil
}

// Save implements PaymentMethodRepository
func (r *PaymentMethodRepositoryImpl) Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	currentTime := time.Now()
	newPaymentMethod.Created_At = currentTime
	result := r.Db.Create(&newPaymentMethod)
	helper.ErrorPanic(result.Error)
	return newPaymentMethod, nil
}

// Update implements PaymentMethodRepository
func (r *PaymentMethodRepositoryImpl) Update(updatedPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	result := r.Db.Model(&model.PaymentMethod{}).Where("id = ?", updatedPaymentMethod.ID).Updates(updatedPaymentMethod)
	helper.ErrorPanic(result.Error)
	return updatedPaymentMethod, nil
}

// FindByName implements PaymentMethodRepository
func (r *PaymentMethodRepositoryImpl) FindByName(name string) (model.PaymentMethod, error) {
	var bor model.PaymentMethod
	result := r.Db.First(&bor, "name = ?", name)

	if result.Error != nil {
		return bor, errors.New("invalid name")
	}
	return bor, nil
}
