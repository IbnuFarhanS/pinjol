package controller

import (
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/pinjol/data/response"
<<<<<<< HEAD
	"github.com/IbnuFarhanS/pinjol/helper"
=======
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
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
<<<<<<< HEAD
	var paymeth model.PaymentMethod
	if err := ctx.ShouldBindJSON(&paymeth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := c.paymentMethodsService.Save(paymeth)
=======
	createpm := model.PaymentMethod{}
	err := ctx.ShouldBindJSON(&createpm)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

<<<<<<< HEAD
=======
	result, err := c.paymentMethodsService.Save(createpm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
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
<<<<<<< HEAD
	helper.ErrorPanic(err)

	updateRol := model.PaymentMethod{ID: id}
	err = ctx.ShouldBindJSON(&updateRol)
	helper.ErrorPanic(err)

	updatedPaymentMethods, err := c.paymentMethodsService.Update(updateRol)
	helper.ErrorPanic(err)
=======
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
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9

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
<<<<<<< HEAD
	helper.ErrorPanic(err)

	c.paymentMethodsService.Delete(id)
	helper.ErrorPanic(err)
=======
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.paymentMethodsService.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9

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
<<<<<<< HEAD
	helper.ErrorPanic(err)
=======
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
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
<<<<<<< HEAD
	helper.ErrorPanic(err)

	pm, err := c.paymentMethodsService.FindById(id)
	helper.ErrorPanic(err)
=======
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pm, err := c.paymentMethodsService.FindById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched PaymentMethods!",
		Data:    pm,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *PaymentMethodsController) FindByName(ctx *gin.Context) {
<<<<<<< HEAD
	userParam := ctx.Param("name")

	pm, err := c.paymentMethodsService.FindByName(userParam)
	helper.ErrorPanic(err)
=======
	paymetParam := ctx.Param("name")

	pm, err := c.paymentMethodsService.FindByName(paymetParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched PaymentMethods!",
		Data:    pm,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
