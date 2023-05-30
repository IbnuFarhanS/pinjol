package service

// import (
// 	"testing"

// 	"github.com/IbnuFarhanS/pinjol/model"
// )

// type mockAcceptStatusRepository struct{}

// func (m *mockAcceptStatusRepository) Save(acceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
// 	// Simulate successful save
// 	return acceptStatus, nil
// }

// func (m *mockAcceptStatusRepository) Delete(id uint) (model.AcceptStatus, error) {
// 	// Simulate successful delete
// 	return model.AcceptStatus{}, nil
// }

// func (m *mockAcceptStatusRepository) FindAll() ([]model.AcceptStatus, error) {
// 	// Simulate finding all accept statuses
// 	acceptStatuses := []model.AcceptStatus{
// 		{ID: 1, TransactionID: 1, Status: true},
// 		{ID: 2, TransactionID: 2, Status: false},
// 	}
// 	return acceptStatuses, nil
// }

// func (m *mockAcceptStatusRepository) FindById(id uint) (model.AcceptStatus, error) {
// 	// Simulate finding an accept status by ID
// 	acceptStatus := model.AcceptStatus{
// 		ID:            uint(id),
// 		TransactionID: 1,
// 		Status:        true,
// 	}
// 	return acceptStatus, nil
// }

// func (m *mockAcceptStatusRepository) Update(acceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
// 	// Simulate successful update
// 	return acceptStatus, nil
// }

// func TestSaveAcceptStatus(t *testing.T) {
// 	service := NewAcceptStatusServiceImpl(&mockAcceptStatusRepository{})

// 	// Test case 1: Valid accept status
// 	acceptStatus := model.AcceptStatus{
// 		TransactionID: 1,
// 		Status:        true,
// 	}
// 	_, err := service.Save(acceptStatus)
// 	if err != nil {
// 		t.Errorf("Expected no error, but got: %v", err)
// 	}

// 	// Test case 2: Invalid accept status (transactionsID is 0)
// 	invalidAcceptStatus := model.AcceptStatus{}
// 	_, err = service.Save(invalidAcceptStatus)
// 	if err == nil {
// 		t.Error("Expected an error, but got none")
// 	} else {
// 		expectedErrorMsg := "id_transaction is required"
// 		if err.Error() != expectedErrorMsg {
// 			t.Errorf("Expected error message: '%s', but got: '%s'", expectedErrorMsg, err.Error())
// 		}
// 	}
// }

// func TestDeleteAcceptStatus(t *testing.T) {
// 	service := NewAcceptStatusServiceImpl(&mockAcceptStatusRepository{})

// 	// Test case: Delete an accept status by ID
// 	id := int64(1)
// 	_, err := service.Delete(uint(id))
// 	if err != nil {
// 		t.Errorf("Expected no error, but got: %v", err)
// 	}
// }

// func TestFindAllAcceptStatuses(t *testing.T) {
// 	service := NewAcceptStatusServiceImpl(&mockAcceptStatusRepository{})

// 	// Test case: Find all accept statuses
// 	acceptStatuses, err := service.FindAll()
// 	if err != nil {
// 		t.Errorf("Expected no error, but got: %v", err)
// 	}

// 	// Check the number of returned accept statuses
// 	expectedCount := 2
// 	if len(acceptStatuses) != expectedCount {
// 		t.Errorf("Expected %d accept statuses, but got: %d", expectedCount, len(acceptStatuses))
// 	}
// }

// func TestFindAcceptStatusByID(t *testing.T) {
// 	service := NewAcceptStatusServiceImpl(&mockAcceptStatusRepository{})

// 	// Test case: Find an accept status by ID
// 	id := int64(1)
// 	acceptStatus, err := service.FindById(uint(id))
// 	if err != nil {
// 		t.Errorf("Expected no error, but got: %v", err)
// 	}

// 	// Check the ID of the returned accept status
// 	if acceptStatus.ID != uint(id) {
// 		t.Errorf("Expected accept status with ID %d, but got: %d", id, acceptStatus.ID)
// 	}
// }

// func TestUpdateAcceptStatus(t *testing.T) {
// 	service := NewAcceptStatusServiceImpl(&mockAcceptStatusRepository{})

// 	// Test case 1: Valid accept status
// 	acceptStatus := model.AcceptStatus{
// 		ID:            1,
// 		TransactionID: 1,
// 		Status:        true,
// 	}
// 	_, err := service.Update(acceptStatus)
// 	if err != nil {
// 		t.Errorf("Expected no error, but got: %v", err)
// 	}

// 	// Test case 2: Invalid accept status (transactionsID is 0)
// 	invalidAcceptStatus := model.AcceptStatus{
// 		ID:            1,
// 		TransactionID: 0,
// 		Status:        false,
// 	}
// 	_, err = service.Update(invalidAcceptStatus)
// 	if err == nil {
// 		t.Error("Expected an error, but got none")
// 	} else {
// 		expectedErrorMsg := "id_transaction tidak boleh kosong"
// 		if err.Error() != expectedErrorMsg {
// 			t.Errorf("Expected error message: '%s', but got: '%s'", expectedErrorMsg, err.Error())
// 		}
// 	}
// }
