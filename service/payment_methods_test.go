package service

import (
	"testing"

	"github.com/IbnuFarhanS/pinjol/model"
)

type mockPaymentMethodRepository struct{}

func (m *mockPaymentMethodRepository) Save(paymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	// Simulate successful save
	return paymentMethod, nil
}

func (m *mockPaymentMethodRepository) Delete(id uint) (model.PaymentMethod, error) {
	// Simulate successful delete
	return model.PaymentMethod{}, nil
}

func (m *mockPaymentMethodRepository) FindAll() ([]model.PaymentMethod, error) {
	// Simulate finding all payment methods
	paymentMethods := []model.PaymentMethod{
		{ID: 1, Name: "Credit Card"},
		{ID: 2, Name: "Bank Transfer"},
	}
	return paymentMethods, nil
}

func (m *mockPaymentMethodRepository) FindById(id uint) (model.PaymentMethod, error) {
	// Simulate finding a payment method by ID
	paymentMethod := model.PaymentMethod{
		ID:   uint(id),
		Name: "Credit Card",
	}
	return paymentMethod, nil
}

func (m *mockPaymentMethodRepository) FindByName(name string) (model.PaymentMethod, error) {
	// Simulate finding a payment method by name
	paymentMethod := model.PaymentMethod{
		ID:   1,
		Name: name,
	}
	return paymentMethod, nil
}

func (m *mockPaymentMethodRepository) Update(paymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	// Simulate successful update
	return paymentMethod, nil
}

func TestSavePaymentMethod(t *testing.T) {

	service := NewPaymentMethodServiceImpl(&mockPaymentMethodRepository{})

	// Test case 1: Valid payment method
	paymentMethod := model.PaymentMethod{
		Name: "Credit Card",
	}
	_, err := service.Save(paymentMethod)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Test case 2: Invalid payment method (name is empty)
	invalidPaymentMethod := model.PaymentMethod{}
	_, err = service.Save(invalidPaymentMethod)
	if err == nil {
		t.Error("Expected an error, but got none")
	} else {
		expectedErrorMsg := "name is required"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message: '%s', but got: '%s'", expectedErrorMsg, err.Error())
		}
	}
}

func TestDeletePaymentMethod(t *testing.T) {
	service := NewPaymentMethodServiceImpl(&mockPaymentMethodRepository{})

	// Test case: Delete a payment method by ID
	id := int64(1)
	_, err := service.Delete(uint(id))
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestFindAllPaymentMethods(t *testing.T) {
	service := NewPaymentMethodServiceImpl(&mockPaymentMethodRepository{})

	// Test case: Find all payment methods
	paymentMethods, err := service.FindAll()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check the number of returned payment methods
	expectedCount := 2
	if len(paymentMethods) != expectedCount {
		t.Errorf("Expected %d payment methods, but got: %d", expectedCount, len(paymentMethods))
	}
}

func TestFindPaymentMethodByID(t *testing.T) {
	service := NewPaymentMethodServiceImpl(&mockPaymentMethodRepository{})

	// Test case: Find a payment method by ID
	id := int64(1)
	paymentMethod, err := service.FindById(uint(id))
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check the ID of the returned payment method
	if paymentMethod.ID != uint(id) {
		t.Errorf("Expected payment method with ID %d, but got: %d", id, paymentMethod.ID)
	}
}

func TestFindPaymentMethodByName(t *testing.T) {
	service := NewPaymentMethodServiceImpl(&mockPaymentMethodRepository{})

	// Test case: Find a payment method by name
	name := "Credit Card"
	paymentMethod, err := service.FindByName(name)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check the name of the returned payment method
	if paymentMethod.Name != name {
		t.Errorf("Expected payment method with name '%s', but got: '%s'", name, paymentMethod.Name)
	}
}

func TestUpdatePaymentMethod(t *testing.T) {

	service := NewPaymentMethodServiceImpl(&mockPaymentMethodRepository{})

	// Test case 1: Valid payment method
	paymentMethod := model.PaymentMethod{
		ID:   1,
		Name: "New Credit Card",
	}
	_, err := service.Update(paymentMethod)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Test case 2: Invalid payment method (name is empty)
	invalidPaymentMethod := model.PaymentMethod{
		ID:   1,
		Name: "",
	}
	_, err = service.Update(invalidPaymentMethod)
	if err == nil {
		t.Error("Expected an error, but got none")
	} else {
		expectedErrorMsg := "name is required"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message: '%s', but got: '%s'", expectedErrorMsg, err.Error())
		}
	}
}
