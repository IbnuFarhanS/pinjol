package service

import "github.com/IbnuFarhanS/pinjol/model"

type PaymentService interface {
	Save(newPayment model.Payment) (model.Payment, error)
	Update(updatePayment model.Payment) (model.Payment, error)
	Delete(id uint) (model.Payment, error)
	FindById(id uint) (model.Payment, error)
	FindAll() ([]model.Payment, error)
}
