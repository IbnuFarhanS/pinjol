package service

import (
	"errors"
	"fmt"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

type PaymentServiceImpl struct {
	PaymentRepository     repository.PaymentRepository
	UserRepository        repository.UserRepository
	TransactionRepository repository.TransactionRepository
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
	transaction, err := s.TransactionRepository.FindById(newPayment.Transaction.ID)
	if err != nil {
		return model.Payment{}, err
	}

	// Periksa apakah pembayaran sudah melebihi jumlah yang harus dibayarkan pada transaksi
	if newPayment.PaymentAmount > transaction.Amount {
		return model.Payment{}, errors.New("payment amount exceeds transaction amount")
	}

	// Perbarui limit pada pengguna terkait
	user, err := s.UserRepository.FindById(transaction.UserID)
	if err != nil {
		return model.Payment{}, err
	}

	user.Limit += newPayment.PaymentAmount

	_, err = s.UserRepository.Update(user)
	if err != nil {
		return model.Payment{}, err
	}

	// Simpan pembayaran
	return s.PaymentRepository.Save(newPayment)

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

func NewPaymentServiceImpl(PaymentRepository repository.PaymentRepository) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository: PaymentRepository,
	}
}
