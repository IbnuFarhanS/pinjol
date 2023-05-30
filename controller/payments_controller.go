package controller

import (
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
type PaymentsController struct {
	paymentsService service.PaymentsService
}

func NewPaymentsController(service service.PaymentsService) *PaymentsController {
	return &PaymentsController{paymentsService: service}
}

func (c *PaymentsController) Insert(ctx *gin.Context) {
	createp := model.Payments{}
	err := ctx.ShouldBindJSON(&createp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.paymentsService.Save(createp)
=======
type PaymentController struct {
	PaymentService service.PaymentService
}

func NewPaymentController(service service.PaymentService) *PaymentController {
	return &PaymentController{PaymentService: service}
}

func (c *PaymentController) Insert(ctx *gin.Context) {
	createp := model.Payment{}
	err := ctx.ShouldBindJSON(&createp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON: " + err.Error()})
		return
	}

	if createp.NextInstallment == 0 {
		webResponse := response.Response{
			Code:    http.StatusOK,
			Status:  "Ok",
			Message: "LUNAS",
			Data:    nil,
		}
		ctx.JSON(http.StatusOK, webResponse)
		return
	}

	result, err := c.PaymentService.Save(createp)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully created PaymentMethods!",
=======
		Message: "Successfully created Payment!",
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

<<<<<<< HEAD
func (c *PaymentsController) Update(ctx *gin.Context) {
=======
func (c *PaymentController) Update(ctx *gin.Context) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
<<<<<<< HEAD
	updatepay := model.Payments{ID: id}
=======
	updatepay := model.Payment{ID: uint(id)}
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	err = ctx.ShouldBindJSON(&updatepay)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

<<<<<<< HEAD
	updatedPayments, err := c.paymentsService.Update(updatepay)
=======
	updatedPayment, err := c.PaymentService.Update(updatepay)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated PaymentMethods!",
<<<<<<< HEAD
		Data:    updatedPayments,
=======
		Data:    updatedPayment,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}

	ctx.JSON(http.StatusOK, webResponse)
}
<<<<<<< HEAD
func (c *PaymentsController) Delete(ctx *gin.Context) {
=======
func (c *PaymentController) Delete(ctx *gin.Context) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

<<<<<<< HEAD
	result, err := c.paymentsService.Delete(id)
=======
	result, err := c.PaymentService.Delete(uint(id))
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
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

<<<<<<< HEAD
func (c *PaymentsController) FindAll(ctx *gin.Context) {
	pm, err := c.paymentsService.FindAll()
=======
func (c *PaymentController) FindAll(ctx *gin.Context) {
	pm, err := c.PaymentService.FindAll()
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
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

<<<<<<< HEAD
func (c *PaymentsController) FindByID(ctx *gin.Context) {
=======
func (c *PaymentController) FindByID(ctx *gin.Context) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

<<<<<<< HEAD
	pay, err := c.paymentsService.FindById(id)
=======
	pay, err := c.PaymentService.FindById(uint(id))
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
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
