package service

import (
	"errors"
<<<<<<< HEAD
=======
	"fmt"
	"time"
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

<<<<<<< HEAD
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
=======
type PaymentServiceImpl struct {
	PaymentRepository       repository.PaymentRepository
	UserRepository          repository.UserRepository
	TransactionRepository   repository.TransactionRepository
	PaymentMethodRepository repository.PaymentMethodRepository
	ProductRepository       repository.ProductRepository
}

// Delete implements BorrowerService
func (s *PaymentServiceImpl) Delete(id uint) (model.Payment, error) {
	return s.PaymentRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *PaymentServiceImpl) FindAll() ([]model.Payment, error) {
	Payment, err := s.PaymentRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for i := range Payment {
		transaction, err := s.TransactionRepository.FindById(Payment[i].TransactionID)
		if err != nil {
			// Tangani kesalahan
			fmt.Println("Error:", err)
			continue
		}

		Payment[i].Transaction = transaction

		Payment[i].NextInstallment = transaction.Total - Payment[i].PaymentAmount
	}

	return Payment, nil
}

// FindById implements BorrowerService
func (s *PaymentServiceImpl) FindById(id uint) (model.Payment, error) {
	return s.PaymentRepository.FindById(id)
}

// Save implements BorrowerService
func (s *PaymentServiceImpl) Save(newPayment model.Payment) (model.Payment, error) {
	tra, err := s.TransactionRepository.FindById(newPayment.TransactionID)
	if err != nil {
		return model.Payment{}, err
	}

	pro, err := s.ProductRepository.FindById(uint(tra.ProductID))
	if err != nil {
		return model.Payment{}, err
	}

	TotalTax := (pro.Interest * tra.Amount) / 100
	Total := tra.Amount + TotalTax
	TotalMonth := Total / float64(pro.Installment)

	if newPayment.PaymentAmount > TotalMonth {
		return model.Payment{}, errors.New("payment amount exceeds transaction total")
	}

	// Update the limit of the related user
	user, err := s.UserRepository.FindById(tra.UserID)
	if err != nil {
		return model.Payment{}, err
	}

	// Save the payment
	newPayment.PaymentDate = time.Now()
	newPayment.NextInstallment = Total - newPayment.PaymentAmount

	payment, err := s.PaymentRepository.Save(newPayment)
	if err != nil {
		return model.Payment{}, err
	}

	user.Limit += (newPayment.PaymentAmount - TotalTax)

	_, err = s.UserRepository.Update(user)
	if err != nil {
		return model.Payment{}, err
	}

	return payment, nil
}

// Update implements BorrowerService
func (s *PaymentServiceImpl) Update(updatePayment model.Payment) (model.Payment, error) {

	var pay model.Payment
	payment_date := pay.PaymentDate

	newPay := model.Payment{
		ID:              updatePayment.ID,
		TransactionID:   updatePayment.TransactionID,
		Transaction:     updatePayment.Transaction,
		PaymentAmount:   updatePayment.PaymentAmount,
		PaymentDate:     payment_date,
		PaymentMethodID: updatePayment.PaymentMethodID,
		PaymentMethod:   model.PaymentMethod{},
	}

	return s.PaymentRepository.Update(newPay)
}

func NewPaymentServiceImpl(
	PaymentRepository repository.PaymentRepository,
	TransactionRepository repository.TransactionRepository,
	UserRepository repository.UserRepository,
	ProductRepository repository.ProductRepository,
) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository:     PaymentRepository,
		TransactionRepository: TransactionRepository,
		UserRepository:        UserRepository,
		ProductRepository:     ProductRepository,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}
}
