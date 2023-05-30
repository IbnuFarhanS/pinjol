package service

import (
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

type PaymentMethodServiceImpl struct {
	PaymentMethodRepository repository.PaymentMethodRepository
}

// Delete implements BorrowerService
func (s *PaymentMethodServiceImpl) Delete(id uint) (model.PaymentMethod, error) {
	return s.PaymentMethodRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *PaymentMethodServiceImpl) FindAll() ([]model.PaymentMethod, error) {
	return s.PaymentMethodRepository.FindAll()
}

// FindById implements BorrowerService
func (s *PaymentMethodServiceImpl) FindById(id uint) (model.PaymentMethod, error) {
	return s.PaymentMethodRepository.FindById(id)
}

// FindByUsername implements BorrowerService
func (s *PaymentMethodServiceImpl) FindByName(name string) (model.PaymentMethod, error) {
	return s.PaymentMethodRepository.FindByName(name)
}

// Save implements BorrowerService
func (s *PaymentMethodServiceImpl) Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {

	newPm := model.PaymentMethod{
		Name:       newPaymentMethod.Name,
		CreatedAt: newPaymentMethod.CreatedAt,
	}
	return s.PaymentMethodRepository.Save(newPm)

}

// Update implements BorrowerService
func (s *PaymentMethodServiceImpl) Update(updatePaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {

	var pm model.PaymentMethod
	create_at := pm.CreatedAt

	newPm := model.PaymentMethod{
		ID:         updatePaymentMethod.ID,
		Name:       updatePaymentMethod.Name,
		CreatedAt: create_at,
	}

	return s.PaymentMethodRepository.Update(newPm)
}

func NewPaymentMethodServiceImpl(paymentMethodRepository repository.PaymentMethodRepository) PaymentMethodService {
	return &PaymentMethodServiceImpl{
		PaymentMethodRepository: paymentMethodRepository,
	}
}
