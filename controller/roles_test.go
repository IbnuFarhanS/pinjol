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

func TestRolesController_Insert(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	RolesService := &mockRoleService{}
	RolesController := NewRoleController(RolesService)

	// Define the test route and bind the controller method to it
	router.POST("/payment-methods", RolesController.Insert)

	// Create a test request body
	requestBody := model.Role{
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
	assert.Equal(t, "Successfully created Roles!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestRolesController_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	RolesService := &mockRoleService{}
	RolesController := NewRoleController(RolesService)

	// Define the test route and bind the controller method to it
	router.PUT("/payment-methods/:id", RolesController.Update)

	// Create a test request body
	requestBody := model.Role{
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
	assert.Equal(t, "Successfully updated Roles!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestRolesController_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	RolesService := &mockRoleService{}
	RolesController := NewRoleController(RolesService)

	// Define the test route and bind the controller method to it
	router.DELETE("/payment-methods/:id", RolesController.Delete)

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
	assert.Equal(t, "Successfully deleted Roles!", response.Message)
	assert.Nil(t, response.Data)

	// Assert any other necessary conditions
}

func TestRolesController_FindAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	RolesService := &mockRoleService{}
	RolesController := NewRoleController(RolesService)

	// Define the test route and bind the controller method to it
	router.GET("/payment-methods", RolesController.FindAll)

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
	assert.Equal(t, "Successfully fetch all Roles data!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestRolesController_FindByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	RolesService := &mockRoleService{}
	RolesController := NewRoleController(RolesService)

	// Define the test route and bind the controller method to it
	router.GET("/payment-methods/:id", RolesController.FindByID)

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
	assert.Equal(t, "Successfully fetched Roles!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestRolesController_FindByName(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	RolesService := &mockRoleService{}
	RolesController := NewRoleController(RolesService)

	// Define the test route and bind the controller method to it
	router.GET("/payment-methods/:name", RolesController.FindByName)

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
	assert.Equal(t, "Successfully fetched Roles!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

// Create a mock service for RoleService
type mockRoleService struct{}

func (m *mockRoleService) Save(Role model.Role) (model.Role, error) {
	// Implement the save method for testing
	return model.Role{}, nil
}

func (m *mockRoleService) Update(Role model.Role) (model.Role, error) {
	// Implement the update method for testing
	return model.Role{}, nil
}

func (m *mockRoleService) Delete(id uint) (model.Role, error) {
	// Implement the delete method for testing
	return model.Role{}, nil
}

func (m *mockRoleService) FindAll() ([]model.Role, error) {
	// Implement the findAll method for testing
	return []model.Role{}, nil
}

func (m *mockRoleService) FindById(id uint) (model.Role, error) {
	// Implement the findById method for testing
	return model.Role{}, nil
}

func (m *mockRoleService) FindByName(name string) (model.Role, error) {
	// Implement the findByName method for testing
	return model.Role{}, nil
}
