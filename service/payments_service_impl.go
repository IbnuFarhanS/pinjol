package service

import (
	"errors"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

type PaymentsServiceImpl struct {
	PaymentsRepository repository.PaymentsRepository
}

// Delete implements BorrowerService
func (s *PaymentsServiceImpl) Delete(id int64) (model.Payments, error) {
	return s.PaymentsRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *PaymentsServiceImpl) FindAll() ([]model.Payments, error) {
	return s.PaymentsRepository.FindAll()
}

// FindById implements BorrowerService
func (s *PaymentsServiceImpl) FindById(id int64) (model.Payments, error) {
	return s.PaymentsRepository.FindById(id)
}

// Save implements BorrowerService
func (s *PaymentsServiceImpl) Save(newPayments model.Payments) (model.Payments, error) {
	// Validate id_transaction
	if newPayments.TransactionsID == 0 {
		return model.Payments{}, errors.New("id transaction is required")
	}
	// Validate id_payment_method
	if newPayments.PaymentMethodID == 0 {
		return model.Payments{}, errors.New("id payment method is required")
	}
	// Validate payment amount
	if newPayments.Payment_Amount == 0 {
		return model.Payments{}, errors.New("payment amount is required")
	}

	newPay := model.Payments{
		TransactionsID:  newPayments.TransactionsID,
		PaymentMethodID: newPayments.PaymentMethodID,
		Payment_Amount:  newPayments.Payment_Amount,
		Payment_Date:    newPayments.Payment_Date,
	}
	return s.PaymentsRepository.Save(newPay)

}

// Update implements BorrowerService
func (s *PaymentsServiceImpl) Update(updatePayments model.Payments) (model.Payments, error) {
	// Validate id_transaction
	if updatePayments.TransactionsID == 0 {
		return model.Payments{}, errors.New("id transaction is required")
	}
	// Validate id_payment_method
	if updatePayments.PaymentMethodID == 0 {
		return model.Payments{}, errors.New("id payment method is required")
	}
	// Validate payment amount
	if updatePayments.Payment_Amount == 0 {
		return model.Payments{}, errors.New("payment amount is required")
	}

	var pay model.Payments
	payment_date := pay.Payment_Date

	newPay := model.Payments{
		ID:              updatePayments.ID,
		TransactionsID:  updatePayments.TransactionsID,
		PaymentMethodID: updatePayments.PaymentMethodID,
		Payment_Amount:  updatePayments.Payment_Amount,
		Payment_Date:    payment_date,
	}

	return s.PaymentsRepository.Update(newPay)
}

func NewPaymentsServiceImpl(PaymentsRepository repository.PaymentsRepository) PaymentsService {
	return &PaymentsServiceImpl{
		PaymentsRepository: PaymentsRepository,
	}
}
