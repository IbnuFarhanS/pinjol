package controller

import (
	"net/http"

	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/gin-gonic/gin"
)

type UploadFileKTPController struct {
	fileKTPService service.UploadFileKTPService
}

func NewUploadFileKTPController(fileKTPService service.UploadFileKTPService) *UploadFileKTPController {
	return &UploadFileKTPController{
		fileKTPService: fileKTPService,
	}
}

func (c *UploadFileKTPController) UploadFileKTP(ctx *gin.Context) {
	file, err := ctx.FormFile("ktp_image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileKTP, err := c.fileKTPService.UploadFileKTP(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Upload Successfully!",
		"filename": fileKTP.Filename,
		"path":     fileKTP.Path})
}
