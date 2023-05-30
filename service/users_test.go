package service

import (
	"testing"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/stretchr/testify/assert"
)

type mockUsersRepository struct{}

func (m *mockUsersRepository) Save(users model.Users) (model.Users, error) {
	// Simulate successful save
	return users, nil
}

func (m *mockUsersRepository) Delete(id int64) (model.Users, error) {
	// Simulate successful delete
	return model.Users{}, nil
}

func (m *mockUsersRepository) FindAll() ([]model.Users, error) {
	// Simulate finding all users
	users := []model.Users{
		{ID: 1, Username: "user1", Name: "User 1"},
		{ID: 2, Username: "user2", Name: "User 2"},
	}
	return users, nil
}

func (m *mockUsersRepository) FindById(id int64) (model.Users, error) {
	// Simulate finding a user by ID
	user := model.Users{
		ID:       id,
		Username: "user1",
		Name:     "User 1",
	}
	return user, nil
}

func (m *mockUsersRepository) FindByUsername(username string) (model.Users, error) {
	// Simulate finding a user by username
	if username == "existinguser" {
		user := model.Users{
			ID:       1,
			Username: username,
			Name:     "Existing User",
		}
		return user, nil
	}
	return model.Users{}, nil
}

func (m *mockUsersRepository) Update(users model.Users) (model.Users, error) {
	// Simulate successful update
	return users, nil
}

func TestSaveUser(t *testing.T) {
	repo := &mockUsersRepository{}
	service := NewUsersServiceImpl(repo)

	t.Run("ValidUser", func(t *testing.T) {
		user := model.Users{
			Username:     "newuser",
			Password:     "password",
			Nik:          "1234567890",
			Name:         "New User",
			Alamat:       "Jalan Raya",
			Phone_Number: "123456789",
			Limit:        2000000,
			RolesID:      1,
		}

		savedUser, err := service.Save(user)
		assert.NoError(t, err)
		assert.Equal(t, user, savedUser)
	})

	t.Run("DuplicateUsername", func(t *testing.T) {
		user := model.Users{
			Username:     "existinguser",
			Password:     "password",
			Nik:          "1234567890",
			Name:         "New User",
			Alamat:       "Jalan Raya",
			Phone_Number: "123456789",
			Limit:        2000000,
			RolesID:      1,
		}

		_, err := service.Save(user)
		assert.EqualError(t, err, "username is already in use")
	})

	t.Run("MissingFields", func(t *testing.T) {
		user := model.Users{
			Username:     "newuser",
			Password:     "password",
			Nik:          "", // Missing nik field
			Name:         "New User",
			Alamat:       "Jalan Raya",
			Phone_Number: "", // Missing phone number field
			Limit:        2000000,
			RolesID:      1,
		}

		_, err := service.Save(user)
		assert.EqualError(t, err, "nik is required")
	})
}

func TestDeleteUser(t *testing.T) {
	service := NewUsersServiceImpl(&mockUsersRepository{})

	// Test case: Delete a user by ID
	id := int64(1)
	_, err := service.Delete(id)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestFindAllUsers(t *testing.T) {
	service := NewUsersServiceImpl(&mockUsersRepository{})

	// Test case: Find all users
	users, err := service.FindAll()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check the number of returned users
	expectedCount := 2
	if len(users) != expectedCount {
		t.Errorf("Expected %d users, but got: %d", expectedCount, len(users))
	}
}

func TestFindUserById(t *testing.T) {
	service := NewUsersServiceImpl(&mockUsersRepository{})

	// Test case: Find a user by ID
	id := int64(1)
	user, err := service.FindById(id)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check the ID of the returned user
	if user.ID != id {
		t.Errorf("Expected user with ID %d, but got: %d", id, user.ID)
	}
}

func TestUpdateUser(t *testing.T) {
	service := NewUsersServiceImpl(&mockUsersRepository{})

	// Test case 1: Valid user
	user := model.Users{
		ID:           1,
		Username:     "updated_user",
		Password:     "updated_password",
		Nik:          "updated_nik",
		Name:         "Updated User",
		Alamat:       "Updated Address",
		Phone_Number: "updated_phone",
		Limit:        3000000,
		RolesID:      2,
	}
	_, err := service.Update(user)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Test case 2: Invalid user (empty username)
	invalidUser := model.Users{}
	_, err = service.Update(invalidUser)
	if err == nil {
		t.Error("Expected an error, but got none")
	} else {
		expectedErrorMsg := "username is required"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message: '%s', but got: '%s'", expectedErrorMsg, err.Error())
		}
	}
}

// func TestFindByUsername(t *testing.T) {
// 	usersRepository := &mockUsersRepository{}
// 	usersService := NewUsersServiceImpl(usersRepository)

// 	// Test existing user
// 	existingUser, err := usersService.FindByUsername("existinguser")
// 	assert.NoError(t, err)
// 	expectedUser := model.Users{
// 		ID:       1,
// 		Username: "existinguser",
// 		Name:     "Existing User",
// 	}
// 	assert.Equal(t, expectedUser, existingUser)

// 	// Test non-existing user
// 	nonExistingUser, err := usersService.FindByUsername("nonexistinguser")
// 	assert.Error(t, err)
// 	assert.EqualError(t, err, "user not found")
// 	assert.Equal(t, model.Users{}, nonExistingUser)
// }
