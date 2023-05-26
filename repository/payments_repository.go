package repository

import "github.com/IbnuFarhanS/pinjol/model"

type PaymentsRepository interface {
	Save(newPayments model.Payments) (model.Payments, error)
	Update(updatePayments model.Payments) (model.Payments, error)
	Delete(id int64) (model.Payments, error)
	FindById(id int64) (model.Payments, error)
	FindAll() ([]model.Payments, error)
	FindByName(name string) (model.Payments, error)
}
