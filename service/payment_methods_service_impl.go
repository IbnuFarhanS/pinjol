package service

import (
	"github.com/pkg/errors"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/go-playground/validator/v10"
)

type PaymentMethodServiceImpl struct {
	PaymentMethodRepository repository.PaymentMethodRepository
	Validate                *validator.Validate
}

// Delete implements BorrowerService
func (s *PaymentMethodServiceImpl) Delete(id int64) (model.PaymentMethod, error) {
	return s.PaymentMethodRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *PaymentMethodServiceImpl) FindAll() ([]model.PaymentMethod, error) {
	return s.PaymentMethodRepository.FindAll()
}

// FindById implements BorrowerService
func (s *PaymentMethodServiceImpl) FindById(id int64) (model.PaymentMethod, error) {
	return s.PaymentMethodRepository.FindById(id)
}

// FindByUsername implements BorrowerService
func (s *PaymentMethodServiceImpl) FindByName(name string) (model.PaymentMethod, error) {
	return s.PaymentMethodRepository.FindByName(name)
}

// Save implements BorrowerService
func (s *PaymentMethodServiceImpl) Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	// Validate name
	if newPaymentMethod.Name == "" {
		return model.PaymentMethod{}, errors.New("name is required")
	}

	newPm := model.PaymentMethod{
		Name:       newPaymentMethod.Name,
		Created_At: newPaymentMethod.Created_At,
	}
	return s.PaymentMethodRepository.Save(newPm)

}

// Update implements BorrowerService
func (s *PaymentMethodServiceImpl) Update(updatePaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	// Validate name
	if updatePaymentMethod.Name == "" {
		return model.PaymentMethod{}, errors.New("name is required")
	}

	var pm model.PaymentMethod
	create_at := pm.Created_At

	newPm := model.PaymentMethod{
		ID:         updatePaymentMethod.ID,
		Name:       updatePaymentMethod.Name,
		Created_At: create_at,
	}

	return s.PaymentMethodRepository.Update(newPm)
}

func NewPaymentMethodServiceImpl(PaymentMethodRepository repository.PaymentMethodRepository, validate *validator.Validate) PaymentMethodService {
	return &PaymentMethodServiceImpl{
		PaymentMethodRepository: PaymentMethodRepository,
		Validate:                validate,
	}
}
