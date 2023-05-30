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
func (r *PaymentMethodRepositoryImpl) Delete(id uint) (model.PaymentMethod, error) {
	var pm model.PaymentMethod
	result := r.Db.Where("id = ?", id).Delete(&pm)
	helper.ErrorPanic(result.Error)
	return pm, nil
}

// FindAll implements PaymentMethodRepository
func (r *PaymentMethodRepositoryImpl) FindAll() ([]model.PaymentMethod, error) {
	var pm []model.PaymentMethod
	results := r.Db.Find(&pm)
	helper.ErrorPanic(results.Error)
	return pm, nil
}

// FindById implements PaymentMethodRepository
func (r *PaymentMethodRepositoryImpl) FindById(id uint) (model.PaymentMethod, error) {
	var pm model.PaymentMethod
	result := r.Db.Find(&pm, "id = ?", id)
	if result.Error != nil {
		return pm, errors.New("paymentMethod is not found")
	}
	return pm, nil
}

// Save implements PaymentMethodRepository
func (r *PaymentMethodRepositoryImpl) Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	currentTime := time.Now()
	newPaymentMethod.CreatedAt = currentTime
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
	var pm model.PaymentMethod
	result := r.Db.First(&pm, "name = ?", name)

	if result.Error != nil {
		return pm, errors.New("invalid name")
	}
	return pm, nil
}
