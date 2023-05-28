package controller

import (
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/gin-gonic/gin"
)

type AcceptStatusController struct {
	acceptStatusService service.AcceptStatusService
}

func NewAcceptStatusController(service service.AcceptStatusService) *AcceptStatusController {
	return &AcceptStatusController{acceptStatusService: service}
}

func (c *AcceptStatusController) Insert(ctx *gin.Context) {
	createLen := model.AcceptStatus{}
	err := ctx.ShouldBindJSON(&createLen)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.acceptStatusService.Save(createLen)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created AcceptStatus!",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *AcceptStatusController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateacc := model.AcceptStatus{ID: id}
	err = ctx.ShouldBindJSON(&updateacc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAcceptStatus, err := c.acceptStatusService.Update(updateacc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated AcceptStatus!",
		Data:    updatedAcceptStatus,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *AcceptStatusController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.acceptStatusService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted AcceptStatus!",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *AcceptStatusController) FindAll(ctx *gin.Context) {
	acc, err := c.acceptStatusService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all AcceptStatus data!",
		Data:    acc,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *AcceptStatusController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	acc, err := c.acceptStatusService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched AcceptStatus!",
		Data:    acc,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
