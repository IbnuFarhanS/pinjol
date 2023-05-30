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
<<<<<<< HEAD
func (r *PaymentMethodRepositoryImpl) Delete(id int64) (model.PaymentMethod, error) {
=======
func (r *PaymentMethodRepositoryImpl) Delete(id uint) (model.PaymentMethod, error) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
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
<<<<<<< HEAD
func (r *PaymentMethodRepositoryImpl) FindById(id int64) (model.PaymentMethod, error) {
=======
func (r *PaymentMethodRepositoryImpl) FindById(id uint) (model.PaymentMethod, error) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
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
<<<<<<< HEAD
	newPaymentMethod.Created_At = currentTime
=======
	newPaymentMethod.CreatedAt = currentTime
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	result := r.Db.Create(&newPaymentMethod)
	helper.ErrorPanic(result.Error)
	return newPaymentMethod, nil
}

// Update implements PaymentMethodRepository
func (r *PaymentMethodRepositoryImpl) Update(updatedPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
<<<<<<< HEAD
	var rol model.PaymentMethod
	created_at := rol.Created_At

	var updatePayMeth = model.PaymentMethod{
		ID:         updatedPaymentMethod.ID,
		Name:       updatedPaymentMethod.Name,
		Created_At: created_at,
	}
	result := r.Db.Model(&updatedPaymentMethod).Updates(updatePayMeth)
=======
	result := r.Db.Model(&model.PaymentMethod{}).Where("id = ?", updatedPaymentMethod.ID).Updates(updatedPaymentMethod)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
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
