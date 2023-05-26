package repository

import "github.com/IbnuFarhanS/pinjol/model"

type PaymentMethodRepository interface {
	Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	Update(updatePaymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	Delete(id int64) (model.PaymentMethod, error)
	FindById(id int64) (model.PaymentMethod, error)
	FindAll() ([]model.PaymentMethod, error)
	FindByName(name string) (model.PaymentMethod, error)
}
