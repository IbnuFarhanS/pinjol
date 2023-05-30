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

func TestAcceptStatusController_Insert(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	acceptStatusService := &mockAcceptStatusService{}
	acceptStatusController := NewAcceptStatusController(acceptStatusService)

	// Define the test route and bind the controller method to it
	router.POST("/accept-status", acceptStatusController.Insert)

	// Create a test request body
	requestBody := model.AcceptStatus{
		// Define your test data here
	}

	// Convert the request body to JSON
	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a test request with the JSON body
	req, _ := http.NewRequest("POST", "/accept-status", bytes.NewBuffer(requestBodyJSON))
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
	assert.Equal(t, "Successfully created AcceptStatus!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestAcceptStatusController_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	acceptStatusService := &mockAcceptStatusService{}
	acceptStatusController := NewAcceptStatusController(acceptStatusService)

	// Define the test route and bind the controller method to it
	router.PUT("/accept-status/:id", acceptStatusController.Update)

	// Create a test request body
	requestBody := model.AcceptStatus{
		ID: 1,
		// Define your test data here
	}

	// Convert the request body to JSON
	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a test request with the JSON body
	req, _ := http.NewRequest("PUT", "/accept-status/1", bytes.NewBuffer(requestBodyJSON))
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
	assert.Equal(t, "Successfully updated AcceptStatus!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestAcceptStatusController_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	acceptStatusService := &mockAcceptStatusService{}
	acceptStatusController := NewAcceptStatusController(acceptStatusService)

	// Define the test route and bind the controller method to it
	router.DELETE("/accept-status/:id", acceptStatusController.Delete)

	// Create a test request
	req, _ := http.NewRequest("DELETE", "/accept-status/1", nil)

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
	assert.Equal(t, "Successfully deleted AcceptStatus!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestAcceptStatusController_FindAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	acceptStatusService := &mockAcceptStatusService{}
	acceptStatusController := NewAcceptStatusController(acceptStatusService)

	// Define the test route and bind the controller method to it
	router.GET("/accept-status", acceptStatusController.FindAll)

	// Create a test request
	req, _ := http.NewRequest("GET", "/accept-status", nil)

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
	assert.Equal(t, "Successfully fetch all AcceptStatus data!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestAcceptStatusController_FindByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	acceptStatusService := &mockAcceptStatusService{}
	acceptStatusController := NewAcceptStatusController(acceptStatusService)

	// Define the test route and bind the controller method to it
	router.GET("/accept-status/:id", acceptStatusController.FindByID)

	// Create a test request
	req, _ := http.NewRequest("GET", "/accept-status/1", nil)

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
	assert.Equal(t, "Successfully fetched AcceptStatus!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

type mockAcceptStatusService struct{}

func (m *mockAcceptStatusService) Save(createAcc model.AcceptStatus) (model.AcceptStatus, error) {
	// Implement your mock save method and return the desired result
	return model.AcceptStatus{}, nil
}

func (m *mockAcceptStatusService) Update(updatedAcc model.AcceptStatus) (model.AcceptStatus, error) {
	// Implement your mock update method and return the desired result
	return model.AcceptStatus{}, nil
}

func (m *mockAcceptStatusService) Delete(id uint) (model.AcceptStatus, error) {
	// Implement your mock delete method and return the desired result
	return model.AcceptStatus{}, nil
}

func (m *mockAcceptStatusService) FindAll() ([]model.AcceptStatus, error) {
	// Implement your mock find all method and return the desired result
	return []model.AcceptStatus{}, nil
}

func (m *mockAcceptStatusService) FindById(id uint) (model.AcceptStatus, error) {
	// Implement your mock find by ID method and return the desired result
	return model.AcceptStatus{}, nil
}
