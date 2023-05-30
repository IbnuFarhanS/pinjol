package service

import "github.com/IbnuFarhanS/pinjol/model"

type AcceptStatusService interface {
	Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error)
	Update(updateAcceptStatus model.AcceptStatus) (model.AcceptStatus, error)
	Delete(id uint) (model.AcceptStatus, error)
	FindById(id uint) (model.AcceptStatus, error)
	FindAll() ([]model.AcceptStatus, error)
}
