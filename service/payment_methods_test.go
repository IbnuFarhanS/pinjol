package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/stretchr/testify/assert"
)

type mockPaymentMethodRepository struct{}

func (m *mockPaymentMethodRepository) Delete(id uint) (model.PaymentMethod, error) {
	// Simulate deleting a payment method
	if id == 1 {
		paymentMethod := model.PaymentMethod{
			ID:        1,
			Name:      "Payment Method 1",
			CreatedAt: time.Now(),
		}
		return paymentMethod, nil
	}
	return model.PaymentMethod{}, errors.New("payment method not found")
}

func (m *mockPaymentMethodRepository) FindAll() ([]model.PaymentMethod, error) {
	// Simulate finding all payment methods
	paymentMethods := []model.PaymentMethod{
		{
			ID:        1,
			Name:      "Payment Method 1",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Payment Method 2",
			CreatedAt: time.Now(),
		},
	}
	return paymentMethods, nil
}

func (m *mockPaymentMethodRepository) FindById(id uint) (model.PaymentMethod, error) {
	// Simulate finding a payment method by ID
	if id == 1 {
		paymentMethod := model.PaymentMethod{
			ID:        1,
			Name:      "Payment Method 1",
			CreatedAt: time.Now(),
		}
		return paymentMethod, nil
	}
	return model.PaymentMethod{}, errors.New("payment method not found")
}

func (m *mockPaymentMethodRepository) FindByName(name string) (model.PaymentMethod, error) {
	// Simulate finding a payment method by name
	if name == "Payment Method 1" {
		paymentMethod := model.PaymentMethod{
			ID:        1,
			Name:      "Payment Method 1",
			CreatedAt: time.Now(),
		}
		return paymentMethod, nil
	}
	return model.PaymentMethod{}, errors.New("payment method not found")
}

func (m *mockPaymentMethodRepository) Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	// Simulate saving a new payment method
	paymentMethod := model.PaymentMethod{
		Name:      newPaymentMethod.Name,
		CreatedAt: newPaymentMethod.CreatedAt,
	}
	return paymentMethod, nil
}

func (m *mockPaymentMethodRepository) Update(updatePaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	// Simulate updating a payment method
	paymentMethod := model.PaymentMethod{
		ID:        updatePaymentMethod.ID,
		Name:      updatePaymentMethod.Name,
		CreatedAt: updatePaymentMethod.CreatedAt,
	}
	return paymentMethod, nil
}

func TestPaymentMethodService(t *testing.T) {
	repo := &mockPaymentMethodRepository{}
	paymentMethodService := service.NewPaymentMethodServiceImpl(repo)

	t.Run("Delete_ValidPaymentMethod", func(t *testing.T) {
		id := uint(1)

		paymentMethod, err := paymentMethodService.Delete(id)
		assert.NoError(t, err)
		assert.Equal(t, id, paymentMethod.ID)
	})

	t.Run("Delete_InvalidPaymentMethod", func(t *testing.T) {
		id := uint(2)

		_, err := paymentMethodService.Delete(id)
		assert.Error(t, err)
		assert.EqualError(t, err, "payment method not found")
	})

	t.Run("FindAll", func(t *testing.T) {
		paymentMethods, err := paymentMethodService.FindAll()
		assert.NoError(t, err)
		assert.Len(t, paymentMethods, 2)
	})

	t.Run("FindById_ValidPaymentMethod", func(t *testing.T) {
		id := uint(1)

		paymentMethod, err := paymentMethodService.FindById(id)
		assert.NoError(t, err)
		assert.Equal(t, id, paymentMethod.ID)
	})

	t.Run("FindById_InvalidPaymentMethod", func(t *testing.T) {
		id := uint(2)

		_, err := paymentMethodService.FindById(id)
		assert.Error(t, err)
		assert.EqualError(t, err, "payment method not found")
	})

	t.Run("FindByName_ValidPaymentMethod", func(t *testing.T) {
		name := "Payment Method 1"

		paymentMethod, err := paymentMethodService.FindByName(name)
		assert.NoError(t, err)
		assert.Equal(t, name, paymentMethod.Name)
	})

	t.Run("FindByName_InvalidPaymentMethod", func(t *testing.T) {
		name := "Payment Method 3"

		_, err := paymentMethodService.FindByName(name)
		assert.Error(t, err)
		assert.EqualError(t, err, "payment method not found")
	})

	t.Run("Save", func(t *testing.T) {
		newPaymentMethod := model.PaymentMethod{
			Name:      "New Payment Method",
			CreatedAt: time.Now(),
		}

		paymentMethod, err := paymentMethodService.Save(newPaymentMethod)
		assert.NoError(t, err)
		assert.Equal(t, newPaymentMethod.Name, paymentMethod.Name)
	})

	t.Run("Update", func(t *testing.T) {
		updatePaymentMethod := model.PaymentMethod{
			ID:        1,
			Name:      "Updated Payment Method",
			CreatedAt: time.Now(),
		}

		paymentMethod, err := paymentMethodService.Update(updatePaymentMethod)
		assert.NoError(t, err)
		assert.Equal(t, updatePaymentMethod.Name, paymentMethod.Name)
	})
}
