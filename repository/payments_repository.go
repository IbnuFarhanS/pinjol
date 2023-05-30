package repository

import "github.com/IbnuFarhanS/pinjol/model"

<<<<<<< HEAD
type PaymentsRepository interface {
	Save(newPayments model.Payments) (model.Payments, error)
	Update(updatePayments model.Payments) (model.Payments, error)
	Delete(id int64) (model.Payments, error)
	FindById(id int64) (model.Payments, error)
	FindAll() ([]model.Payments, error)
=======
type PaymentRepository interface {
	Save(newPayment model.Payment) (model.Payment, error)
	Update(updatePayment model.Payment) (model.Payment, error)
	Delete(id uint) (model.Payment, error)
	FindById(id uint) (model.Payment, error)
	FindAll() ([]model.Payment, error)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
