package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTransactionService struct {
	mock.Mock
}

func (m *mockTransactionService) Save(Transaction model.Transaction, userid uint) (model.Transaction, error) {
	args := m.Called(Transaction)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func (m *mockTransactionService) Update(Transaction model.Transaction) (model.Transaction, error) {
	args := m.Called(Transaction)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func (m *mockTransactionService) Delete(id uint) (model.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func (m *mockTransactionService) FindAll() ([]model.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]model.Transaction), args.Error(1)
}

func (m *mockTransactionService) FindById(id uint) (model.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func (m *mockTransactionService) FindByName(name string) (model.Transaction, error) {
	args := m.Called(name)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func TestTransactionController_Delete(t *testing.T) {
	transactionID := uint(1)

	// Mock the transaction service
	mockService := &mockTransactionService{}
	mockService.On("Delete", transactionID).Return(model.Transaction{}, nil)

	controller := NewTransactionController(mockService)

	// Create a test router using gin
	router := gin.Default()
	router.DELETE("/transactions/:id", controller.Delete)

	// Create a test request
	request, _ := http.NewRequest("DELETE", "/transactions/"+strconv.Itoa(int(transactionID)), nil)

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
	assert.Equal(t, response.Message, "Successfully deleted Transactions!")
	assert.Nil(t, response.Data)
}

func TestTransactionController_FindAll(t *testing.T) {
	// Mock the transaction service
	mockService := &mockTransactionService{}
	mockService.On("FindAll").Return([]model.Transaction{}, nil)

	controller := NewTransactionController(mockService)

	// Create a test router using gin
	router := gin.Default()
	router.GET("/transactions", controller.FindAll)

	// Create a test request
	request, _ := http.NewRequest("GET", "/transactions", nil)

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
	assert.Equal(t, response.Message, "Successfully fetch all Transactions data!")
	assert.Equal(t, response.Data, []model.Transaction{})
}

func TestTransactionController_FindByID(t *testing.T) {
	transactionID := uint(1)
	transaction := model.Transaction{
		ID: transactionID,
		// Set the transaction fields
	}

	// Mock the transaction service
	mockService := &mockTransactionService{}
	mockService.On("FindById", transactionID).Return(transaction, nil)

	controller := NewTransactionController(mockService)

	// Create a test router using gin
	router := gin.Default()
	router.GET("/transactions/:id", controller.FindByID)

	// Create a test request
	request, _ := http.NewRequest("GET", "/transactions/"+strconv.Itoa(int(transactionID)), nil)

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
	assert.Equal(t, response.Message, "Successfully fetched Transactions!")
	assert.Equal(t, response.Data, transaction)
}

func TestTransactionController_Insert(t *testing.T) {
	// Mock the transaction service
	mockService := &mockTransactionService{}
	mockService.On("Save", mock.Anything).Return(model.Transaction{}, nil)

	controller := NewTransactionController(mockService)

	// Create a test router using gin
	router := gin.Default()
	router.POST("/transactions", controller.Insert)

	// Create a test request body
	requestBody := model.Transaction{
		// Set the transaction fields
	}
	requestJSON, _ := json.Marshal(requestBody)

	// Create a test request
	request, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(requestJSON))
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
	assert.Equal(t, response.Message, "Successfully created Transactions!")
	assert.NotNil(t, response.Data)
}

func TestTransactionController_Update(t *testing.T) {
	transactionID := uint(1)
	transaction := model.Transaction{
		ID: transactionID,
		// Set the transaction fields
	}

	// Mock the transaction service
	mockService := &mockTransactionService{}
	mockService.On("Update", transaction).Return(transaction, nil)

	controller := NewTransactionController(mockService)

	// Create a test router using gin
	router := gin.Default()
	router.PUT("/transactions/:id", controller.Update)

	// Create a test request body
	requestBody := model.Transaction{
		ID: transactionID,
		// Set the transaction fields
	}
	requestJSON, _ := json.Marshal(requestBody)

	// Create a test request
	request, _ := http.NewRequest("PUT", "/transactions/"+strconv.Itoa(int(transactionID)), bytes.NewBuffer(requestJSON))
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
	assert.Equal(t, response.Message, "Successfully updated Transactions!")
	assert.Equal(t, response.Data, transaction)
}

func TestTransactionController_ExportToCSV(t *testing.T) {
	// Mock the transaction service
	mockService := &mockTransactionService{}
	mockService.On("FindAll").Return([]model.Transaction{}, nil)

	controller := NewTransactionController(mockService)

	// Create a test router using gin
	router := gin.Default()
	router.GET("/transactions/export/csv", controller.ExportToCSV)

	// Create a test request
	request, _ := http.NewRequest("GET", "/transactions/export/csv", nil)

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
	assert.Equal(t, response.Message, "Transactions exported to CSV successfully")
	assert.Nil(t, response.Data)
}
