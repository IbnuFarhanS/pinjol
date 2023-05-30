package service

import (
	"errors"
	"fmt"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/go-playground/validator/v10"
)

type PaymentsServiceImpl struct {
	PaymentsRepository     repository.PaymentsRepository
	UsersRepository        repository.UsersRepository
	TransactionsRepository repository.TransactionsRepository
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

	for i := range payments {
		transaction, err := s.TransactionsRepository.FindById(payments[i].TransactionsID)
		if err != nil {
			// Tangani kesalahan
			fmt.Println("Error:", err)
			continue
		}

		payments[i].Transactions = transaction

		payments[i].NextInstallment = transaction.Total - payments[i].Payment_Amount
	}

	return payments, nil
}

// FindById implements BorrowerService
func (s *PaymentsServiceImpl) FindById(id int64) (model.Payments, error) {
	return s.PaymentsRepository.FindById(id)
}

// Save implements BorrowerService
func (s *PaymentsServiceImpl) Save(newPayments model.Payments) (model.Payments, error) {
	transaction, err := s.TransactionsRepository.FindById(newPayments.Transactions.ID)
	if err != nil {
		return model.Payments{}, err
	}

	// Periksa apakah pembayaran sudah melebihi jumlah yang harus dibayarkan pada transaksi
	if newPayments.Payment_Amount > transaction.Amount {
		return model.Payments{}, errors.New("payment amount exceeds transaction amount")
	}

	// Perbarui limit pada pengguna terkait
	user, err := s.UsersRepository.FindById(transaction.UsersID)
	if err != nil {
		return model.Payments{}, err
	}

	user.Limit += newPayments.Payment_Amount

	_, err = s.UsersRepository.Update(user)
	if err != nil {
		return model.Payments{}, err
	}

	// Simpan pembayaran
	return s.PaymentsRepository.Save(newPayments)

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
