package controller

import (
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/gin-gonic/gin"
)

type PaymentsController struct {
	paymentsService     service.PaymentsService
	transactionsService service.TransactionsService
}

func NewPaymentsController(paymentsService service.PaymentsService, transactionsService service.TransactionsService) *PaymentsController {
	return &PaymentsController{
		paymentsService:     paymentsService,
		transactionsService: transactionsService,
	}
}

func (c *PaymentsController) Insert(ctx *gin.Context) {
	createp := model.Payments{}
	err := ctx.ShouldBindJSON(&createp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// traid := createp.TransactionsID
	// transaction, err := c.transactionsService.FindById(traid)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	result, err := c.paymentsService.Save(createp)
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

func (c *PaymentsController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatepay := model.Payments{ID: id}
	err = ctx.ShouldBindJSON(&updatepay)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPayments, err := c.paymentsService.Update(updatepay)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.paymentsService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted PaymentMethods!",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *PaymentsController) FindAll(ctx *gin.Context) {
	pm, err := c.paymentsService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func (c *PaymentsController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pay, err := c.paymentsService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched PaymentMethods!",
		Data:    pay,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
