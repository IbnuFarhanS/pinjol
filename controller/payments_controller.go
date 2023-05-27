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

type PaymentsController struct {
	paymentsService service.PaymentsService
}

func NewPaymentsController(service service.PaymentsService) *PaymentsController {
	return &PaymentsController{paymentsService: service}
}

func (c *PaymentsController) Insert(ctx *gin.Context) {
	createp := model.Payments{}
	err := ctx.ShouldBindJSON(&createp)
	helper.ErrorPanic(err)

	c.paymentsService.Save(createp)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created PaymentMethods!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *PaymentsController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	updatepay := model.Payments{ID: id}
	err = ctx.ShouldBindJSON(&updatepay)
	helper.ErrorPanic(err)

	updatedPayments, err := c.paymentsService.Update(updatepay)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated PaymentMethods!",
		Data:    updatedPayments,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *PaymentsController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	c.paymentsService.Delete(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted PaymentMethods!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *PaymentsController) FindAll(ctx *gin.Context) {
	pm, err := c.paymentsService.FindAll()
	helper.ErrorPanic(err)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all PaymentMethods data!",
		Data:    pm,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *PaymentsController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	pay, err := c.paymentsService.FindById(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched PaymentMethods!",
		Data:    pay,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
