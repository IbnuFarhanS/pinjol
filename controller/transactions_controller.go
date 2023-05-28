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

type TransactionsController struct {
	transactionsService service.TransactionsService
}

func NewTransactionsController(service service.TransactionsService) *TransactionsController {
	return &TransactionsController{transactionsService: service}
}

func (c *TransactionsController) Insert(ctx *gin.Context) {
	tra := model.Transactions{}
	err := ctx.ShouldBindJSON(&tra)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	currentUserID, _ := ctx.Get("currentUserID")
	userID, _ := currentUserID.(int64)

	result, err := c.transactionsService.Save(tra, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created Transactions!",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *TransactionsController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	updatetra := model.Transactions{ID: id}
	err = ctx.ShouldBindJSON(&updatetra)
	helper.ErrorPanic(err)

	updatedTransactions, err := c.transactionsService.Update(updatetra)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated Transactions!",
		Data:    updatedTransactions,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *TransactionsController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	c.transactionsService.Delete(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted Transactions!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *TransactionsController) FindAll(ctx *gin.Context) {
	len, err := c.transactionsService.FindAll()
	helper.ErrorPanic(err)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all Transactions data!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *TransactionsController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	tra, err := c.transactionsService.FindById(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Transactions!",
		Data:    tra,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
