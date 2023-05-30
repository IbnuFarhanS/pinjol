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

func TestProductController_Insert(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	productService := &mockProductService{}
	productController := NewProductController(productService)

	// Define the test route and bind the controller method to it
	router.POST("/products", productController.Insert)

	// Create a test request body
	requestBody := model.Product{
		// Define your test data here
	}

	// Convert the request body to JSON
	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a test request with the JSON body
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(requestBodyJSON))
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

func TestProductController_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	productService := &mockProductService{}
	productController := NewProductController(productService)

	// Define the test route and bind the controller method to it
	router.PUT("/products/:id", productController.Update)

	// Create a test request body
	requestBody := model.Product{
		// Define your test data here
	}

	// Convert the request body to JSON
	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a test request with the JSON body and URL parameter
	req, _ := http.NewRequest("PUT", "/products/1", bytes.NewBuffer(requestBodyJSON))
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

func TestProductController_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	productService := &mockProductService{}
	productController := NewProductController(productService)

	// Define the test route and bind the controller method to it
	router.DELETE("/products/:id", productController.Delete)

	// Create a test request with the URL parameter
	req, _ := http.NewRequest("DELETE", "/products/1", nil)

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

func TestProductController_FindAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	productService := &mockProductService{}
	productController := NewProductController(productService)

	// Define the test route and bind the controller method to it
	router.GET("/products", productController.FindAll)

	// Create a test request
	req, _ := http.NewRequest("GET", "/products", nil)

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

func TestProductController_FindByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	productService := &mockProductService{}
	productController := NewProductController(productService)

	// Define the test route and bind the controller method to it
	router.GET("/products/:id", productController.FindByID)

	// Create a test request with the URL parameter
	req, _ := http.NewRequest("GET", "/products/1", nil)

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

func TestProductController_FindByName(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	productService := &mockProductService{}
	productController := NewProductController(productService)

	// Define the test route and bind the controller method to it
	router.GET("/payment-methods/:name", productController.FindByName)

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

// Define a mock product service for testing
type mockProductService struct{}

func (m *mockProductService) Save(product model.Product) (model.Product, error) {
	// Implement the save method for testing
	return model.Product{}, nil
}

func (m *mockProductService) Update(product model.Product) (model.Product, error) {
	// Implement the update method for testing
	return model.Product{}, nil
}

func (m *mockProductService) Delete(id uint) (model.Product, error) {
	// Implement the delete method for testing
	return model.Product{}, nil
}

func (m *mockProductService) FindAll() ([]model.Product, error) {
	// Implement the findAll method for testing
	return []model.Product{}, nil
}

func (m *mockProductService) FindById(id uint) (model.Product, error) {
	// Implement the findById method for testing
	return model.Product{}, nil
}

func (m *mockProductService) FindByName(name string) (model.Product, error) {
	// Implement the findByName method for testing
	return model.Product{}, nil
}
