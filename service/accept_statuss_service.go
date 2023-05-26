package service

import "github.com/IbnuFarhanS/pinjol/model"

type AcceptStatusService interface {
	Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error)
	Update(updateAcceptStatus model.AcceptStatus) (model.AcceptStatus, error)
	Delete(id int64) (model.AcceptStatus, error)
	FindById(id int64) (model.AcceptStatus, error)
	FindAll() ([]model.AcceptStatus, error)
	FindByName(username string) (model.AcceptStatus, error)
}
