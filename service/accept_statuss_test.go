package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/stretchr/testify/assert"
)

type mockAcceptStatusRepository struct{}

func (m *mockAcceptStatusRepository) Delete(id uint) (model.AcceptStatus, error) {
	// Simulate deleting an accept status
	if id == 1 {
		acceptStatus := model.AcceptStatus{
			ID:        1,
			Status:    true,
			CreatedAt: time.Now(),
		}
		return acceptStatus, nil
	}
	return model.AcceptStatus{}, errors.New("accept status not found")
}

func (m *mockAcceptStatusRepository) FindAll() ([]model.AcceptStatus, error) {
	// Simulate finding all accept statuses
	acceptStatuses := []model.AcceptStatus{
		{
			ID:        1,
			Status:    true,
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Status:    false,
			CreatedAt: time.Now(),
		},
	}
	return acceptStatuses, nil
}

func (m *mockAcceptStatusRepository) FindById(id uint) (model.AcceptStatus, error) {
	// Simulate finding an accept status by ID
	if id == 1 {
		acceptStatus := model.AcceptStatus{
			ID:        1,
			Status:    true,
			CreatedAt: time.Now(),
		}
		return acceptStatus, nil
	}
	return model.AcceptStatus{}, errors.New("accept status not found")
}

func (m *mockAcceptStatusRepository) Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	// Simulate saving a new accept status
	acceptStatus := model.AcceptStatus{
		Status:    newAcceptStatus.Status,
		CreatedAt: newAcceptStatus.CreatedAt,
	}
	return acceptStatus, nil
}

func (m *mockAcceptStatusRepository) Update(updatedAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	// Simulate updating an accept status
	acceptStatus := model.AcceptStatus{
		ID:        updatedAcceptStatus.ID,
		Status:    updatedAcceptStatus.Status,
		CreatedAt: updatedAcceptStatus.CreatedAt,
	}
	return acceptStatus, nil
}

type mockTransactionRepository struct{}

func (m *mockTransactionRepository) UpdateStatus(id uint, status bool) error {
	// Simulate updating transaction status
	return nil
}

func (m *mockTransactionRepository) Delete(id uint) (model.Transaction, error) {
	// Simulate deleting a transaction
	return model.Transaction{}, nil
}

func (m *mockTransactionRepository) FindAll() ([]model.Transaction, error) {
	// Simulate finding all transactions
	transactions := []model.Transaction{
		{
			ID:        1,
			Status:    true,
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Status:    false,
			CreatedAt: time.Now(),
		},
	}
	return transactions, nil
}

func (m *mockTransactionRepository) FindById(id uint) (model.Transaction, error) {
	// Simulate finding a transaction by ID
	if id == 1 {
		transaction := model.Transaction{
			ID:        1,
			Status:    true,
			CreatedAt: time.Now(),
		}
		return transaction, nil
	}
	return model.Transaction{}, errors.New("transaction not found")
}

func (m *mockTransactionRepository) FindByUserID(userID uint) ([]model.Transaction, error) {
	// Simulate finding transactions by user ID
	transactions := []model.Transaction{
		{
			ID:        1,
			Status:    true,
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Status:    false,
			CreatedAt: time.Now(),
		},
	}
	return transactions, nil
}

func (m *mockTransactionRepository) Save(newTransaction model.Transaction) (model.Transaction, error) {
	// Simulate saving a new transaction
	transaction := model.Transaction{
		Status:    newTransaction.Status,
		CreatedAt: newTransaction.CreatedAt,
	}
	return transaction, nil
}

func (m *mockTransactionRepository) Update(updatedTransaction model.Transaction) (model.Transaction, error) {
	// Simulate updating a transaction
	transaction := model.Transaction{
		ID:        updatedTransaction.ID,
		Status:    updatedTransaction.Status,
		CreatedAt: updatedTransaction.CreatedAt,
	}
	return transaction, nil
}

func TestAcceptStatusService(t *testing.T) {
	acceptStatusRepo := &mockAcceptStatusRepository{}
	transactionRepo := &mockTransactionRepository{}
	acceptStatusService := service.NewAcceptStatusServiceImpl(acceptStatusRepo, transactionRepo)

	t.Run("Delete_ValidAcceptStatus", func(t *testing.T) {
		id := uint(1)

		acceptStatus, err := acceptStatusService.Delete(id)
		assert.NoError(t, err)
		assert.Equal(t, id, acceptStatus.ID)
	})

	t.Run("Delete_InvalidAcceptStatus", func(t *testing.T) {
		id := uint(2)

		_, err := acceptStatusService.Delete(id)
		assert.Error(t, err)
		assert.EqualError(t, err, "accept status not found")
	})

	t.Run("FindAll", func(t *testing.T) {
		acceptStatuses, err := acceptStatusService.FindAll()
		assert.NoError(t, err)
		assert.Len(t, acceptStatuses, 2)
	})

	t.Run("FindById_ValidAcceptStatus", func(t *testing.T) {
		id := uint(1)

		acceptStatus, err := acceptStatusService.FindById(id)
		assert.NoError(t, err)
		assert.Equal(t, id, acceptStatus.ID)
	})

	t.Run("FindById_InvalidAcceptStatus", func(t *testing.T) {
		id := uint(2)

		_, err := acceptStatusService.FindById(id)
		assert.Error(t, err)
		assert.EqualError(t, err, "accept status not found")
	})

	t.Run("Save", func(t *testing.T) {
		newAcceptStatus := model.AcceptStatus{
			Status:    true,
			CreatedAt: time.Now(),
		}

		acceptStatus, err := acceptStatusService.Save(newAcceptStatus)
		assert.NoError(t, err)
		assert.Equal(t, newAcceptStatus.Status, acceptStatus.Status)
	})

	t.Run("Update", func(t *testing.T) {
		updateAcceptStatus := model.AcceptStatus{
			ID:        1,
			Status:    false,
			CreatedAt: time.Now(),
		}

		acceptStatus, err := acceptStatusService.Update(updateAcceptStatus)
		assert.NoError(t, err)
		assert.Equal(t, updateAcceptStatus.Status, acceptStatus.Status)
	})
}
