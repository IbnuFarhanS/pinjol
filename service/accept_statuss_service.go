package service

import "github.com/IbnuFarhanS/pinjol/model"

type AcceptStatusService interface {
	Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error)
	Update(updateAcceptStatus model.AcceptStatus) (model.AcceptStatus, error)
<<<<<<< HEAD
	Delete(id int64) (model.AcceptStatus, error)
	FindById(id int64) (model.AcceptStatus, error)
=======
	Delete(id uint) (model.AcceptStatus, error)
	FindById(id uint) (model.AcceptStatus, error)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	FindAll() ([]model.AcceptStatus, error)
}
