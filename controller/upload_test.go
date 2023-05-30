package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockUploadFileKTPService struct {
	UploadFileKTPFunc func(fileHeader *multipart.FileHeader) (model.FileKTP, error)
}

func (m *mockUploadFileKTPService) UploadFileKTP(fileHeader *multipart.FileHeader) (model.FileKTP, error) {
	return m.UploadFileKTPFunc(fileHeader)
}

func TestUploadFileKTPController_UploadFileKTP(t *testing.T) {
	fileHeader := &multipart.FileHeader{
		Filename: "test.jpg",
	}

	expectedFileKTP := model.FileKTP{
		Filename: "test.jpg",
		Path:     "/path/to/test.jpg",
	}

	// Mock the upload file KTP service
	mockService := &mockUploadFileKTPService{}
	mockService.UploadFileKTPFunc = func(fileHeader *multipart.FileHeader) (model.FileKTP, error) {
		return expectedFileKTP, nil
	}

	controller := NewUploadFileKTPController(mockService)

	// Create a test router using gin
	router := gin.Default()
	router.POST("/upload-file-ktp", controller.UploadFileKTP)

	// Create a test multipart form request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("ktp_image", fileHeader.Filename)
	file, _ := fileHeader.Open()
	io.Copy(part, file)
	writer.Close()

	// Create a test request
	request, _ := http.NewRequest("POST", "/upload-file-ktp", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(recorder, request)

	// Verify the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body
	var response map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verify the response data
	assert.Equal(t, response["message"], "Upload Successfully!")
	assert.Equal(t, response["filename"], expectedFileKTP.Filename)
	assert.Equal(t, response["path"], expectedFileKTP.Path)
}

func TestUploadFileKTPController_UploadFileKTP_Error(t *testing.T) {
	fileHeader := &multipart.FileHeader{
		Filename: "test.jpg",
	}

	expectedError := errors.New("Upload failed")

	// Mock the upload file KTP service
	mockService := &mockUploadFileKTPService{}
	mockService.UploadFileKTPFunc = func(fileHeader *multipart.FileHeader) (model.FileKTP, error) {
		return model.FileKTP{}, expectedError
	}

	controller := NewUploadFileKTPController(mockService)

	// Create a test router using gin
	router := gin.Default()
	router.POST("/upload-file-ktp", controller.UploadFileKTP)

	// Create a test multipart form request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("ktp_image", fileHeader.Filename)
	file, _ := fileHeader.Open()
	io.Copy(part, file)
	writer.Close()

	// Create a test request
	request, _ := http.NewRequest("POST", "/upload-file-ktp", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(recorder, request)

	// Verify the response
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	// Parse the response body
	var response map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verify the response data
	assert.Equal(t, response["error"], expectedError.Error())
}
