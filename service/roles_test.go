package service

import (
	"testing"

	"github.com/IbnuFarhanS/pinjol/model"
)

type mockRolesRepository struct{}

func (m *mockRolesRepository) Save(roles model.Roles) (model.Roles, error) {
	// Simulate successful save
	return roles, nil
}

func (m *mockRolesRepository) Delete(id int64) (model.Roles, error) {
	// Simulate successful delete
	return model.Roles{}, nil
}

func (m *mockRolesRepository) FindAll() ([]model.Roles, error) {
	// Simulate finding all roles
	roles := []model.Roles{
		{ID: 1, Name: "User"},
		{ID: 2, Name: "Admin"},
	}
	return roles, nil
}

func (m *mockRolesRepository) FindById(id int64) (model.Roles, error) {
	// Simulate finding a roles by ID
	roles := model.Roles{
		ID:   id,
		Name: "User",
	}
	return roles, nil
}

func (m *mockRolesRepository) FindByName(name string) (model.Roles, error) {
	// Simulate finding a roles by name
	roles := model.Roles{
		ID:   1,
		Name: name,
	}
	return roles, nil
}

func (m *mockRolesRepository) Update(roles model.Roles) (model.Roles, error) {
	// Simulate successful update
	return roles, nil
}

func TestSaveRoles(t *testing.T) {

	service := NewRolesServiceImpl(&mockRolesRepository{})

	// Test case 1: Valid roles
	roles := model.Roles{
		Name: "User",
	}
	_, err := service.Save(roles)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Test case 2: Invalid roles (name is empty)
	invalidRoles := model.Roles{}
	_, err = service.Save(invalidRoles)
	if err == nil {
		t.Error("Expected an error, but got none")
	} else {
		expectedErrorMsg := "name is required"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message: '%s', but got: '%s'", expectedErrorMsg, err.Error())
		}
	}
}

func TestDeleteRoles(t *testing.T) {
	service := NewRolesServiceImpl(&mockRolesRepository{})

	// Test case: Delete a roles by ID
	id := int64(1)
	_, err := service.Delete(id)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestFindAllRoles(t *testing.T) {
	service := NewRolesServiceImpl(&mockRolesRepository{})

	// Test case: Find all roles
	roles, err := service.FindAll()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check the number of returned roles
	expectedCount := 2
	if len(roles) != expectedCount {
		t.Errorf("Expected %d payment methods, but got: %d", expectedCount, len(roles))
	}
}

func TestFindRolesByID(t *testing.T) {
	service := NewRolesServiceImpl(&mockRolesRepository{})

	// Test case: Find a roles by ID
	id := int64(1)
	roles, err := service.FindById(id)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check the ID of the returned roles
	if roles.ID != id {
		t.Errorf("Expected roles with ID %d, but got: %d", id, roles.ID)
	}
}

func TestFindRolesByName(t *testing.T) {
	service := NewRolesServiceImpl(&mockRolesRepository{})

	// Test case: Find a roles by name
	name := "User"
	roles, err := service.FindByName(name)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check the name of the returned roles
	if roles.Name != name {
		t.Errorf("Expected roles with name '%s', but got: '%s'", name, roles.Name)
	}
}

func TestUpdateRoles(t *testing.T) {

	service := NewRolesServiceImpl(&mockRolesRepository{})

	// Test case 1: Valid roles
	roles := model.Roles{
		ID:   1,
		Name: "New User",
	}
	_, err := service.Update(roles)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Test case 2: Invalid roles (name is empty)
	invalidRoles := model.Roles{
		ID:   1,
		Name: "",
	}
	_, err = service.Update(invalidRoles)
	if err == nil {
		t.Error("Expected an error, but got none")
	} else {
		expectedErrorMsg := "name is required"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message: '%s', but got: '%s'", expectedErrorMsg, err.Error())
		}
	}
}
