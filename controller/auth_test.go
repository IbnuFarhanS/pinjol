package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IbnuFarhanS/pinjol/data/request"
	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthController_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	authService := &mockAuthService{}
	authController := NewAuthController(authService)

	// Define the test route and bind the controller method to it
	router.POST("/login", authController.Login)

	// Create a test request body
	requestBody := request.LoginRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	// Convert the request body to JSON
	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a test request with the JSON body
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBodyJSON))
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
	assert.Equal(t, "Successfully log in!", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

func TestAuthController_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	authService := &mockAuthService{}
	authController := NewAuthController(authService)

	// Define the test route and bind the controller method to it
	router.POST("/register", authController.Register)

	// Create a test request body
	requestBody := request.CreateUsersRequest{
		// Define your test data here
	}

	// Convert the request body to JSON
	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a test request with the JSON body
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBodyJSON))
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
	assert.Equal(t, "Successfully created user!", response.Message)

	// Assert any other necessary conditions
}

func TestAuthController_FindAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router and controller instance
	router := gin.Default()
	authService := &mockAuthService{}
	authController := NewAuthController(authService)

	// Define the test route and bind the controller method to it
	router.GET("/users", authController.FindAll)

	// Create a test request
	req, _ := http.NewRequest("GET", "/users", nil)

	// Set a test value for currentUser in the context
	ctx := req.Context()
	ctx = context.WithValue(ctx, "currentUser", "testuser")

	// Assign the modified context back to the request
	req = req.WithContext(ctx)

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
	assert.Equal(t, "Users retrieved successfully", response.Message)
	assert.NotNil(t, response.Data)

	// Assert any other necessary conditions
}

// Implement mock service methods for the AuthService interface
type mockAuthService struct{}

func (m *mockAuthService) Login(loginRequest request.LoginRequest) (string, error) {
	// Implement your mock login method and return the desired result
	return "mocked_token", nil
}

func (m *mockAuthService) Register(createUsersRequest request.CreateUsersRequest) {
	// Implement your mock register method and return the desired result
	return
}

func (m *mockAuthService) FindAll() ([]model.User, error) {
	// Implement your mock find all method and return the desired result
	return []model.User{}, nil
}

func (m *mockAuthService) FindByUserID(userID uint) (model.User, error) {
	// Implement your mock find by user ID method and return the desired result
	return model.User{}, nil
}

func (m *mockAuthService) FindByUsername(username string) (model.User, error) {
	// Implement your mock find by username method and return the desired result
	return model.User{}, nil
}
