package service

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/google/uuid"
)

type UploadFileKTPServiceImpl struct {
}

func NewUploadFileKTPService() UploadFileKTPService {
	return &UploadFileKTPServiceImpl{}
}

// UploadFileKTP uploads the KTP file to the local directory and returns the file information.
func (s *UploadFileKTPServiceImpl) UploadFileKTP(file *multipart.FileHeader) (model.FileKTP, error) {
	// Generate a unique filename for the uploaded file
	filename := generateUniqueFilename(file.Filename)

	// Specify the directory to save the uploaded files
	uploadDir := "img/"

	// Create the directory if it doesn't exist
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return model.FileKTP{}, err
	}

	// Generate the file path
	filePath := filepath.Join(uploadDir, filename)

	// Save the uploaded file to the specified path
	err = saveUploadedFile(file, filePath)
	if err != nil {
		return model.FileKTP{}, err
	}

	// Create and return the FileKTP struct
	fileKTP := model.FileKTP{
		Filename: filename,
		Path:     filePath,
	}

	return fileKTP, nil
}

// Helper function to generate a unique filename
func generateUniqueFilename(originalFilename string) string {
	extension := filepath.Ext(originalFilename)
	filename := uuid.New().String() + extension
	return filename
}

// Helper function to save the uploaded file to the specified path
func saveUploadedFile(file *multipart.FileHeader, filePath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}
