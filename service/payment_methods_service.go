package service

import "github.com/IbnuFarhanS/pinjol/model"

type PaymentMethodService interface {
	Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	Update(updatePaymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	Delete(id uint) (model.PaymentMethod, error)
	FindById(id uint) (model.PaymentMethod, error)
	FindAll() ([]model.PaymentMethod, error)
	FindByName(name string) (model.PaymentMethod, error)
}
