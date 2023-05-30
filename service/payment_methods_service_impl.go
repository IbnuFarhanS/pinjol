package service

import (
<<<<<<< HEAD
	"github.com/pkg/errors"

=======
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

type PaymentMethodServiceImpl struct {
	PaymentMethodRepository repository.PaymentMethodRepository
}

<<<<<<< HEAD
// Delete implements PaymentMethod Service
func (s *PaymentMethodServiceImpl) Delete(id int64) (model.PaymentMethod, error) {
	return s.PaymentMethodRepository.Delete(id)
}

// FindAll implements PaymentMethod Service
=======
// Delete implements BorrowerService
func (s *PaymentMethodServiceImpl) Delete(id uint) (model.PaymentMethod, error) {
	return s.PaymentMethodRepository.Delete(id)
}

// FindAll implements BorrowerService
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
func (s *PaymentMethodServiceImpl) FindAll() ([]model.PaymentMethod, error) {
	return s.PaymentMethodRepository.FindAll()
}

<<<<<<< HEAD
// FindById implements PaymentMethod Service
func (s *PaymentMethodServiceImpl) FindById(id int64) (model.PaymentMethod, error) {
	return s.PaymentMethodRepository.FindById(id)
}

// FindByUsername implements PaymentMethod Service
=======
// FindById implements BorrowerService
func (s *PaymentMethodServiceImpl) FindById(id uint) (model.PaymentMethod, error) {
	return s.PaymentMethodRepository.FindById(id)
}

// FindByUsername implements BorrowerService
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
func (s *PaymentMethodServiceImpl) FindByName(name string) (model.PaymentMethod, error) {
	return s.PaymentMethodRepository.FindByName(name)
}

<<<<<<< HEAD
// Save implements PaymentMethod Service
func (s *PaymentMethodServiceImpl) Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	// Validate name
	if newPaymentMethod.Name == "" {
		return model.PaymentMethod{}, errors.New("name is required")
	}

	newPm := model.PaymentMethod{
		Name:       newPaymentMethod.Name,
		Created_At: newPaymentMethod.Created_At,
=======
// Save implements BorrowerService
func (s *PaymentMethodServiceImpl) Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {

	newPm := model.PaymentMethod{
		Name:       newPaymentMethod.Name,
		CreatedAt: newPaymentMethod.CreatedAt,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}
	return s.PaymentMethodRepository.Save(newPm)

}

<<<<<<< HEAD
// Update implements PaymentMethod Service
func (s *PaymentMethodServiceImpl) Update(updatePaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	// Validate name
	if updatePaymentMethod.Name == "" {
		return model.PaymentMethod{}, errors.New("name is required")
	}

	var pm model.PaymentMethod
	create_at := pm.Created_At
=======
// Update implements BorrowerService
func (s *PaymentMethodServiceImpl) Update(updatePaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {

	var pm model.PaymentMethod
	create_at := pm.CreatedAt
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9

	newPm := model.PaymentMethod{
		ID:         updatePaymentMethod.ID,
		Name:       updatePaymentMethod.Name,
<<<<<<< HEAD
		Created_At: create_at,
=======
		CreatedAt: create_at,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}

	return s.PaymentMethodRepository.Update(newPm)
}

<<<<<<< HEAD
func NewPaymentMethodServiceImpl(PaymentMethodRepository repository.PaymentMethodRepository) PaymentMethodService {
	return &PaymentMethodServiceImpl{
		PaymentMethodRepository: PaymentMethodRepository,
=======
func NewPaymentMethodServiceImpl(paymentMethodRepository repository.PaymentMethodRepository) PaymentMethodService {
	return &PaymentMethodServiceImpl{
		PaymentMethodRepository: paymentMethodRepository,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}
}
