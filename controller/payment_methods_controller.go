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

type PaymentMethodsController struct {
	paymentMethodsService service.PaymentMethodService
}

func NewPaymentMethodsController(service service.PaymentMethodService) *PaymentMethodsController {
	return &PaymentMethodsController{paymentMethodsService: service}
}

func (c *PaymentMethodsController) Insert(ctx *gin.Context) {
	var paymeth model.PaymentMethod
	if err := ctx.ShouldBindJSON(&paymeth); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := c.paymentMethodsService.Save(paymeth)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created PaymentMethods!",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *PaymentMethodsController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	updateRol := model.PaymentMethod{ID: id}
	err = ctx.ShouldBindJSON(&updateRol)
	helper.ErrorPanic(err)

	updatedPaymentMethods, err := c.paymentMethodsService.Update(updateRol)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated PaymentMethods!",
		Data:    updatedPaymentMethods,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *PaymentMethodsController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	c.paymentMethodsService.Delete(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted PaymentMethods!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *PaymentMethodsController) FindAll(ctx *gin.Context) {
	pm, err := c.paymentMethodsService.FindAll()
	helper.ErrorPanic(err)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all PaymentMethods data!",
		Data:    pm,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *PaymentMethodsController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	pm, err := c.paymentMethodsService.FindById(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched PaymentMethods!",
		Data:    pm,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *PaymentMethodsController) FindByName(ctx *gin.Context) {
	userParam := ctx.Param("name")

	pm, err := c.paymentMethodsService.FindByName(userParam)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched PaymentMethods!",
		Data:    pm,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
