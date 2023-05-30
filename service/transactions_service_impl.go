package service

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/go-playground/validator/v10"
)

type TransactionsServiceImpl struct {
	TransactionsRepository repository.TransactionsRepository
	UsersRepository        repository.UsersRepository
	ProductsRepository     repository.ProductsRepository
	ProductsService        ProductsService
}

// Delete implements BorrowerService
func (s *TransactionsServiceImpl) Delete(id int64) (model.Transactions, error) {
	return s.TransactionsRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *TransactionsServiceImpl) FindAll() ([]model.Transactions, error) {
	transactions, err := s.TransactionsRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for i := range transactions {
		transactions[i].TotalTax = (transactions[i].Amount * transactions[i].Products.Bunga) / 100
		transactions[i].Total = transactions[i].TotalTax + transactions[i].Amount
	}

	// for i := range transactions {
	// 	product, err := s.ProductsService.FindById(transactions[i].ProductsID)
	// 	if err != nil {
	// 		// Handle error
	// 		fmt.Println("Error:", err)
	// 		continue
	// 	}
	// 	// fmt.Println("iniadalahidproduct", transactions[i].ProductsID)
	// 	transactions[i].Products = product
	// 	transactions[i].TotalTax = (transactions[i].Amount * transactions[i].Products.Bunga) / 100
	// 	transactions[i].Total = transactions[i].TotalTax + transactions[i].Amount
	// }

	return transactions, nil
}

// FindById implements BorrowerService
func (s *TransactionsServiceImpl) FindById(id int64) (model.Transactions, error) {
	return s.TransactionsRepository.FindById(id)
}

// Save implements BorrowerService
func (s *TransactionsServiceImpl) Save(newTransactions model.Transactions, userid int64) (model.Transactions, error) {

	user, err := s.UsersRepository.FindById(userid)
	if err != nil {
		return model.Transactions{}, err
	}

	// Periksa batas pengguna
	if newTransactions.Amount > user.Limit {
		return model.Transactions{}, errors.New("Amount exceeds user's limit")
	}

	// Kurangi batas pengguna dengan jumlah pinjaman
	user.Limit -= newTransactions.Amount

	// Simpan perubahan ke basis data
	_, err = s.UsersRepository.Update(user)
	if err != nil {
		return model.Transactions{}, err
	}

	created_at := time.Now()
	due_date := created_at.AddDate(0, 1, 0)
	newTra := model.Transactions{
		// Products:   newTransactions.Products,
		ProductsID: newTransactions.ProductsID,
		UsersID:    userid,
		Status:     false,
		Amount:     newTransactions.Amount,
		Created_At: created_at,
		Due_Date:   due_date,
	}
	transaction, err := s.TransactionsRepository.Save(newTra)
	if err != nil {
		return model.Transactions{}, err
	}

	return transaction, nil
}

// Update implements BorrowerService
func (s *TransactionsServiceImpl) Update(updateTransactions model.Transactions) (model.Transactions, error) {

	var tra model.Transactions
	create_at := tra.Created_At

	newTra := model.Transactions{
		ID:         updateTransactions.ID,
		Products:   updateTransactions.Products,
		Users:      updateTransactions.Users,
		Status:     updateTransactions.Status,
		Amount:     updateTransactions.Amount,
		Due_Date:   updateTransactions.Due_Date,
		Created_At: create_at,
	}

	return s.TransactionsRepository.Update(newTra)
}

func NewTransactionsServiceImpl(transactionsRepository repository.TransactionsRepository, validate *validator.Validate, usersRepo repository.UsersRepository) TransactionsService {
	return &TransactionsServiceImpl{
		TransactionsRepository: transactionsRepository,
		UsersRepository:        usersRepo,
	}
}
