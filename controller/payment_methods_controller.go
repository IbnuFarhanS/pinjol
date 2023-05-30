package controller

import (
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/pinjol/data/response"
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
	createpm := model.PaymentMethod{}
	err := ctx.ShouldBindJSON(&createpm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.paymentMethodsService.Save(createpm)
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
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateRol := model.PaymentMethod{ID: uint(id)}
	err = ctx.ShouldBindJSON(&updateRol)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedPaymentMethods, err := c.paymentMethodsService.Update(updateRol)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.paymentMethodsService.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pm, err := c.paymentMethodsService.FindById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched PaymentMethods!",
		Data:    pm,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *PaymentMethodsController) FindByName(ctx *gin.Context) {
	paymetParam := ctx.Param("name")

	pm, err := c.paymentMethodsService.FindByName(paymetParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched PaymentMethods!",
		Data:    pm,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
