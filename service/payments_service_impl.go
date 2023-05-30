package service

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/go-playground/validator/v10"
)

type PaymentsServiceImpl struct {
	PaymentsRepository       repository.PaymentsRepository
	UsersRepository          repository.UsersRepository
	TransactionsRepository   repository.TransactionsRepository
	PaymentMethodsRepository repository.PaymentMethodRepository
}

// Delete implements BorrowerService
func (s *PaymentsServiceImpl) Delete(id int64) (model.Payments, error) {
	return s.PaymentsRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *PaymentsServiceImpl) FindAll() ([]model.Payments, error) {
	payments, err := s.PaymentsRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return payments, nil
}

// FindById implements BorrowerService
func (s *PaymentsServiceImpl) FindById(id int64) (model.Payments, error) {
	return s.PaymentsRepository.FindById(id)
}

// Save implements BorrowerService
func (s *PaymentsServiceImpl) Save(newPayments model.Payments) (model.Payments, error) {

	tra, err := s.TransactionsRepository.FindById(newPayments.TransactionsID)
	if err != nil {
		return model.Payments{}, err
	}

	if newPayments.Payment_Amount > tra.Total {
		return model.Payments{}, errors.New("payment amount exceeds transaction total")
	}

	// Perbarui limit pada pengguna terkait
	user, err := s.UsersRepository.FindById(tra.UsersID)
	if err != nil {
		return model.Payments{}, err
	}

	pm, err := s.PaymentMethodsRepository.FindById(newPayments.PaymentMethodID)
	if err != nil {
		return model.Payments{}, err
	}

	// Simpan pembayaran
	newPayments.Payment_Date = time.Now()
	newPayments.NextInstallment = tra.Total - newPayments.Payment_Amount

	newpay := model.Payments{
		TransactionsID:  newPayments.TransactionsID,
		PaymentMethodID: pm.ID,
		Payment_Amount:  newPayments.Payment_Amount,
		Payment_Date:    newPayments.Payment_Date,
		NextInstallment: newPayments.NextInstallment,
	}
	payment, err := s.PaymentsRepository.Save(newpay)
	if err != nil {
		return model.Payments{}, err
	}

	user.Limit += (newPayments.Payment_Amount - tra.TotalTax)

	_, err = s.UsersRepository.Update(user)
	if err != nil {
		return model.Payments{}, err
	}

	return payment, nil

}

// Update implements BorrowerService
func (s *PaymentsServiceImpl) Update(updatePayments model.Payments) (model.Payments, error) {

	var pay model.Payments
	payment_date := pay.Payment_Date

	newPay := model.Payments{
		ID:             updatePayments.ID,
		Transactions:   updatePayments.Transactions,
		Payment_Method: updatePayments.Payment_Method,
		Payment_Amount: updatePayments.Payment_Amount,
		Payment_Date:   payment_date,
	}

	return s.PaymentsRepository.Update(newPay)
}

func NewPaymentsServiceImpl(paymentsRepository repository.PaymentsRepository, validate *validator.Validate) PaymentsService {
	return &PaymentsServiceImpl{
		PaymentsRepository: paymentsRepository,
	}
}
