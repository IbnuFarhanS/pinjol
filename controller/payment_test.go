package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPaymentController_Insert(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	paymentService := &mockPaymentService{}
	paymentController := NewPaymentController(paymentService)

	// Define the test route and bind the controller method to it
	router.POST("/payments", paymentController.Insert)

	// Create a test request body
	requestBody := model.Payment{
		// Define your test data here
	}

	// Convert the request body to JSON
	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a test request with the JSON body
	req, _ := http.NewRequest("POST", "/payments", bytes.NewBuffer(requestBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a test response recorder
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, res.Code)

	// Parse the response body
	var response response.Response
	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Assert the response data or message as needed
	assert.Equal(t, "Successfully created Payment!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestPaymentController_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	paymentService := &mockPaymentService{}
	paymentController := NewPaymentController(paymentService)

	// Define the test route and bind the controller method to it
	router.PUT("/payments/:id", paymentController.Update)

	// Create a test request body
	requestBody := model.Payment{
		// Define your test data here
	}

	// Convert the request body to JSON
	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a test request with the JSON body and URL parameter
	req, _ := http.NewRequest("PUT", "/payments/1", bytes.NewBuffer(requestBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a test response recorder
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, res.Code)

	// Parse the response body
	var response response.Response
	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Assert the response data or message as needed
	assert.Equal(t, "Successfully updated PaymentMethods!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestPaymentController_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	paymentService := &mockPaymentService{}
	paymentController := NewPaymentController(paymentService)

	// Define the test route and bind the controller method to it
	router.DELETE("/payments/:id", paymentController.Delete)

	// Create a test request with the URL parameter
	req, _ := http.NewRequest("DELETE", "/payments/1", nil)

	// Create a test response recorder
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, res.Code)

	// Parse the response body
	var response response.Response
	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Assert the response data or message as needed
	assert.Equal(t, "Successfully deleted PaymentMethods!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestPaymentController_FindAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	paymentService := &mockPaymentService{}
	paymentController := NewPaymentController(paymentService)

	// Define the test route and bind the controller method to it
	router.GET("/payments", paymentController.FindAll)

	// Create a test request
	req, _ := http.NewRequest("GET", "/payments", nil)

	// Create a test response recorder
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, res.Code)

	// Parse the response body
	var response response.Response
	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Assert the response data or message as needed
	assert.Equal(t, "Successfully fetch all PaymentMethods data!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestPaymentController_FindByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	paymentService := &mockPaymentService{}
	paymentController := NewPaymentController(paymentService)

	// Define the test route and bind the controller method to it
	router.GET("/payments/:id", paymentController.FindByID)

	// Create a test request with the URL parameter
	req, _ := http.NewRequest("GET", "/payments/1", nil)

	// Create a test response recorder
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, res.Code)

	// Parse the response body
	var response response.Response
	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Assert the response data or message as needed
	assert.Equal(t, "Successfully fetched PaymentMethods!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

// Define a mock payment service for testing
type mockPaymentService struct{}

func (m *mockPaymentService) Save(payment model.Payment) (model.Payment, error) {
	// Implement the save method for testing
	return model.Payment{}, nil
}

func (m *mockPaymentService) Update(payment model.Payment) (model.Payment, error) {
	// Implement the update method for testing
	return model.Payment{}, nil
}

func (m *mockPaymentService) Delete(id uint) (model.Payment, error) {
	// Implement the delete method for testing
	return model.Payment{}, nil
}

func (m *mockPaymentService) FindAll() ([]model.Payment, error) {
	// Implement the findAll method for testing
	return []model.Payment{}, nil
}

func (m *mockPaymentService) FindById(id uint) (model.Payment, error) {
	// Implement the findById method for testing
	return model.Payment{}, nil
}
