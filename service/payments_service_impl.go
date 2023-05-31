package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

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

		pro, err := s.ProductRepository.FindById(transaction.ProductID)
		if err != nil {
			// Tangani kesalahan
			fmt.Println("Error:", err)
			continue
		}

		totaltax := (transaction.Amount * pro.Interest) / 100
		total := totaltax + transaction.Amount
		totalmonth := total / float64(pro.Installment)
		Payment[i].NextInstallment = total - totalmonth
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

	user.Limit += (newPayment.PaymentAmount - (TotalTax / float64(pro.Installment)))

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
	}
}
