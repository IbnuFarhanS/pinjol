package service

import (
	"mime/multipart"

	"github.com/IbnuFarhanS/pinjol/model"
)

type UploadFileKTPService interface {
	UploadFileKTP(file *multipart.FileHeader) (model.FileKTP, error)
}
