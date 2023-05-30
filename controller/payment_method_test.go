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

func TestPaymentMethodsController_Insert(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	paymentMethodsService := &mockPaymentMethodService{}
	paymentMethodsController := NewPaymentMethodsController(paymentMethodsService)

	// Define the test route and bind the controller method to it
	router.POST("/payment-methods", paymentMethodsController.Insert)

	// Create a test request body
	requestBody := model.PaymentMethod{
		// Define your test data here
	}

	// Convert the request body to JSON
	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a test request with the JSON body
	req, _ := http.NewRequest("POST", "/payment-methods", bytes.NewBuffer(requestBodyJSON))
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
	assert.Equal(t, "Successfully created PaymentMethods!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestPaymentMethodsController_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	paymentMethodsService := &mockPaymentMethodService{}
	paymentMethodsController := NewPaymentMethodsController(paymentMethodsService)

	// Define the test route and bind the controller method to it
	router.PUT("/payment-methods/:id", paymentMethodsController.Update)

	// Create a test request body
	requestBody := model.PaymentMethod{
		// Define your test data here
	}

	// Convert the request body to JSON
	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a test request with the JSON body and URL parameter
	req, _ := http.NewRequest("PUT", "/payment-methods/1", bytes.NewBuffer(requestBodyJSON))
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

func TestPaymentMethodsController_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	paymentMethodsService := &mockPaymentMethodService{}
	paymentMethodsController := NewPaymentMethodsController(paymentMethodsService)

	// Define the test route and bind the controller method to it
	router.DELETE("/payment-methods/:id", paymentMethodsController.Delete)

	// Create a test request with the URL parameter
	req, _ := http.NewRequest("DELETE", "/payment-methods/1", nil)

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
	assert.Nil(t, response.Data)

	// Assert any other necessary conditions
}

func TestPaymentMethodsController_FindAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	paymentMethodsService := &mockPaymentMethodService{}
	paymentMethodsController := NewPaymentMethodsController(paymentMethodsService)

	// Define the test route and bind the controller method to it
	router.GET("/payment-methods", paymentMethodsController.FindAll)

	// Create a test request
	req, _ := http.NewRequest("GET", "/payment-methods", nil)

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

func TestPaymentMethodsController_FindByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	paymentMethodsService := &mockPaymentMethodService{}
	paymentMethodsController := NewPaymentMethodsController(paymentMethodsService)

	// Define the test route and bind the controller method to it
	router.GET("/payment-methods/:id", paymentMethodsController.FindByID)

	// Create a test request with the URL parameter
	req, _ := http.NewRequest("GET", "/payment-methods/1", nil)

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

func TestPaymentMethodsController_FindByName(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	paymentMethodsService := &mockPaymentMethodService{}
	paymentMethodsController := NewPaymentMethodsController(paymentMethodsService)

	// Define the test route and bind the controller method to it
	router.GET("/payment-methods/:name", paymentMethodsController.FindByName)

	// Create a test request with the URL parameter
	req, _ := http.NewRequest("GET", "/payment-methods/example", nil)

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

// Create a mock service for PaymentMethodService
type mockPaymentMethodService struct{}

func (m *mockPaymentMethodService) Save(paymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	// Implement the save method for testing
	return model.PaymentMethod{}, nil
}

func (m *mockPaymentMethodService) Update(paymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	// Implement the update method for testing
	return model.PaymentMethod{}, nil
}

func (m *mockPaymentMethodService) Delete(id uint) (model.PaymentMethod, error) {
	// Implement the delete method for testing
	return model.PaymentMethod{}, nil
}

func (m *mockPaymentMethodService) FindAll() ([]model.PaymentMethod, error) {
	// Implement the findAll method for testing
	return []model.PaymentMethod{}, nil
}

func (m *mockPaymentMethodService) FindById(id uint) (model.PaymentMethod, error) {
	// Implement the findById method for testing
	return model.PaymentMethod{}, nil
}

func (m *mockPaymentMethodService) FindByName(name string) (model.PaymentMethod, error) {
	// Implement the findByName method for testing
	return model.PaymentMethod{}, nil
}
