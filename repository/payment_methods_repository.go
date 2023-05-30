package repository

import "github.com/IbnuFarhanS/pinjol/model"

type PaymentMethodRepository interface {
	Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	Update(updatePaymentMethod model.PaymentMethod) (model.PaymentMethod, error)
<<<<<<< HEAD
	Delete(id int64) (model.PaymentMethod, error)
	FindById(id int64) (model.PaymentMethod, error)
=======
	Delete(id uint) (model.PaymentMethod, error)
	FindById(id uint) (model.PaymentMethod, error)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	FindAll() ([]model.PaymentMethod, error)
	FindByName(name string) (model.PaymentMethod, error)
}
