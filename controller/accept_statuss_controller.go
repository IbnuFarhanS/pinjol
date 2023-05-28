package controller

import (
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/helper"
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
	var accstat model.AcceptStatus
	if err := ctx.ShouldBindJSON(&accstat); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := c.acceptStatusService.Save(accstat)
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
	helper.ErrorPanic(err)

	updateRol := model.AcceptStatus{ID: id}
	err = ctx.ShouldBindJSON(&updateRol)
	helper.ErrorPanic(err)

	updatedAcceptStatus, err := c.acceptStatusService.Update(updateRol)
	helper.ErrorPanic(err)

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
	helper.ErrorPanic(err)

	c.acceptStatusService.Delete(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted AcceptStatus!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *AcceptStatusController) FindAll(ctx *gin.Context) {
	len, err := c.acceptStatusService.FindAll()
	helper.ErrorPanic(err)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all AcceptStatus data!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *AcceptStatusController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	len, err := c.acceptStatusService.FindById(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched AcceptStatus!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
