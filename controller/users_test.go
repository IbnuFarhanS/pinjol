package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockUserService struct {
	SaveFunc        func(newUser model.User) (model.User, error)
	UpdateFunc      func(updatedUser model.User) (model.User, error)
	DeleteFunc      func(id uint) (model.User, error)
	FindAllFunc     func() ([]model.User, error)
	FindByIdFunc    func(id uint) (model.User, error)
	FindByUsernamee func(username string) (model.User, error)
}

func (m *mockUserService) Save(newUser model.User) (model.User, error) {
	return m.SaveFunc(newUser)
}

func (m *mockUserService) Update(updatedUser model.User) (model.User, error) {
	return m.UpdateFunc(updatedUser)
}

func (m *mockUserService) Delete(id uint) (model.User, error) {
	return m.DeleteFunc(id)
}

func (m *mockUserService) FindAll() ([]model.User, error) {
	return m.FindAllFunc()
}

func (m *mockUserService) FindById(id uint) (model.User, error) {
	return m.FindByIdFunc(id)
}

func (m *mockUserService) FindByUsername(username string) (model.User, error) {
	return m.FindByUsername(username)
}

func TestUserController_Insert(t *testing.T) {
	user := model.User{
		// set user fields
	}

	expectedResult := model.User{
		// set expected user fields
	}

	mockService := &mockUserService{}
	mockService.SaveFunc = func(user model.User) (model.User, error) {
		return expectedResult, nil
	}

	controller := NewUserController(mockService)

	// Create a test router using gin
	router := gin.Default()
	router.POST("/users", controller.Insert)

	// Convert the user struct to JSON
	userJSON, _ := json.Marshal(user)

	// Create a test request
	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	request.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(recorder, request)

	// Verify the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body
	var response response.Response
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verify the response data
	assert.Equal(t, response.Code, 200)
	assert.Equal(t, response.Status, "Ok")
	assert.Equal(t, response.Message, "Successfully created User!")
	assert.Equal(t, response.Data, expectedResult)
}

func TestUserController_Insert_Error(t *testing.T) {
	user := model.User{
		// set user fields
	}

	expectedError := errors.New("Save user failed")

	mockService := &mockUserService{}
	mockService.SaveFunc = func(user model.User) (model.User, error) {
		return model.User{}, expectedError
	}

	controller := NewUserController(mockService)

	// Create a test router using gin
	router := gin.Default()
	router.POST("/users", controller.Insert)

	// Convert the user struct to JSON
	userJSON, _ := json.Marshal(user)

	// Create a test request
	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	request.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(recorder, request)

	// Verify the response
	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	// Parse the response body
	var response response.Response
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verify the response data
	assert.Equal(t, response.Code, 400)
	assert.Equal(t, response.Status, "Bad Request")
	assert.Equal(t, response.Message, expectedError.Error())
	assert.Nil(t, response.Data)
}

// Test other UserController methods similarly
