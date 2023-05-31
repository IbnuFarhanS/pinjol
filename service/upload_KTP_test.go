package service_test

import (
	"errors"
	"mime/multipart"
	"testing"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/stretchr/testify/assert"
)

type mockFileKTPRepository struct{}

func (m *mockFileKTPRepository) Save(fileKTP model.FileKTP) (model.FileKTP, error) {
	// Simulate saving a file KTP
	fileKTP.Path = "/path/to/file.jpg"
	return fileKTP, nil
}

type mockFileUploader struct{}

func (m *mockFileUploader) Upload(file *multipart.FileHeader) (string, error) {
	// Simulate uploading a file
	return "/path/to/file.jpg", nil
}

type mockInvalidFileUploader struct{}

func (m *mockInvalidFileUploader) Upload(file *multipart.FileHeader) (string, error) {
	// Simulate uploading a file
	return "", errors.New("upload failed")
}

func TestUploadFileKTPService(t *testing.T) {
	uploadService := service.NewUploadFileKTPService()

	t.Run("UploadFileKTP_Success", func(t *testing.T) {
		file := &multipart.FileHeader{
			Filename: "ktp.jpg",
		}

		uploadService.UploadFileKTP(file)
		// assert.NoError(t, err)
		// assert.Equal(t, "", fileKTP.Path)
	})

	t.Run("UploadFileKTP_InvalidFile", func(t *testing.T) {
		file := &multipart.FileHeader{
			Filename: "invalid.txt",
		}

		_, err := uploadService.UploadFileKTP(file)
		assert.Error(t, err)
	})

	t.Run("UploadFileKTP_UploadFailed", func(t *testing.T) {
		file := &multipart.FileHeader{
			Filename: "ktp.jpg",
		}

		_, err := uploadService.UploadFileKTP(file)
		assert.Error(t, err)
	})
}
