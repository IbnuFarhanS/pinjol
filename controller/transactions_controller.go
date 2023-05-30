package controller

import (
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
<<<<<<< HEAD
	"github.com/gin-gonic/gin"
)

type TransactionsController struct {
	transactionsService service.TransactionsService
}

func NewTransactionsController(service service.TransactionsService) *TransactionsController {
	return &TransactionsController{transactionsService: service}
}

func (c *TransactionsController) Insert(ctx *gin.Context) {
	createtra := model.Transactions{}
=======
	"github.com/IbnuFarhanS/pinjol/utils"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController(service service.TransactionService) *TransactionController {
	return &TransactionController{transactionService: service}
}

func (c *TransactionController) Insert(ctx *gin.Context) {
	createtra := model.Transaction{}
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	err := ctx.ShouldBindJSON(&createtra)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	currentUserID, _ := ctx.Get("currentUserID")
	userID, _ := currentUserID.(int64)

<<<<<<< HEAD
	result, err := c.transactionsService.Save(createtra, userID)
=======
	result, err := c.transactionService.Save(createtra, uint(userID))
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
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

<<<<<<< HEAD
func (c *TransactionsController) Update(ctx *gin.Context) {
=======
func (c *TransactionController) Update(ctx *gin.Context) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

<<<<<<< HEAD
	updatetra := model.Transactions{ID: id}
=======
	updatetra := model.Transaction{ID: uint(id)}
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	err = ctx.ShouldBindJSON(&updatetra)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

<<<<<<< HEAD
	updatedTransactions, err := c.transactionsService.Update(updatetra)
=======
	updatedTransactions, err := c.transactionService.Update(updatetra)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated Transactions!",
		Data:    updatedTransactions,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
<<<<<<< HEAD
func (c *TransactionsController) Delete(ctx *gin.Context) {
=======
func (c *TransactionController) Delete(ctx *gin.Context) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

<<<<<<< HEAD
	c.transactionsService.Delete(id)
=======
	c.transactionService.Delete(uint(id))
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted Transactions!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

<<<<<<< HEAD
func (c *TransactionsController) FindAll(ctx *gin.Context) {
	len, err := c.transactionsService.FindAll()
=======
func (c *TransactionController) FindAll(ctx *gin.Context) {
	len, err := c.transactionService.FindAll()
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all Transactions data!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

<<<<<<< HEAD
func (c *TransactionsController) FindByID(ctx *gin.Context) {
=======
func (c *TransactionController) FindByID(ctx *gin.Context) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

<<<<<<< HEAD
	tra, err := c.transactionsService.FindById(id)
=======
	tra, err := c.transactionService.FindById(uint(id))
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Transactions!",
		Data:    tra,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
<<<<<<< HEAD
=======

func (c *TransactionController) ExportToCSV(ctx *gin.Context) {
	transactions, err := c.transactionService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	filePath := "export/transactions.csv" // Ganti dengan jalur file CSV yang diinginkan

	err = utils.ExportTransactionsToCSV(transactions, filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transactions exported to CSV successfully"})
}

func (controller *TransactionController) FindAllTransactions(ctx *gin.Context) {
	currentUser := ctx.GetInt64("currentUserID")
	transactions, err := controller.transactionService.FindAll()
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	filteredTra := make([]model.Transaction, 0)
	for _, transaction := range transactions {
		if int64(transaction.UserID) == currentUser {
			filteredTra = append(filteredTra, transaction)
		}
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Transactions retrieved successfully",
		Data:    filteredTra,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
