package controller

import (
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/IbnuFarhanS/pinjol/utils"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService service.TransactionService
}

func NewTransactionController(service service.TransactionService) *TransactionController {
	return &TransactionController{TransactionService: service}
}

func (c *TransactionController) Insert(ctx *gin.Context) {
	createtra := model.Transaction{}
	err := ctx.ShouldBindJSON(&createtra)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	currentUserID, _ := ctx.Get("currentUserID")
	userID, _ := currentUserID.(int64)

	result, err := c.TransactionService.Save(createtra, uint(userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created Transaction!",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *TransactionController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatetra := model.Transaction{ID: uint(id)}
	err = ctx.ShouldBindJSON(&updatetra)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatedTransaction, err := c.TransactionService.Update(updatetra)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated Transaction!",
		Data:    updatedTransaction,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *TransactionController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.TransactionService.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted Transaction!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *TransactionController) FindAll(ctx *gin.Context) {
	len, err := c.TransactionService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all Transaction data!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *TransactionController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	tra, err := c.TransactionService.FindById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Transaction!",
		Data:    tra,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *TransactionController) ExportToCSV(ctx *gin.Context) {
	Transaction, err := c.TransactionService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	filePath := "export/Transaction.csv" // Ganti dengan jalur file CSV yang diinginkan

	err = utils.ExportTransactionsToCSV(Transaction, filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction exported to CSV successfully"})
}

func (controller *TransactionController) FindAllTransaction(ctx *gin.Context) {
	currentUser := ctx.GetUint("currentUserID")
	Transaction, err := controller.TransactionService.FindAll()
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
	for i := range Transaction {
		if Transaction[i].UserID == currentUser {
			// fmt.Println("WADADWADWADAWD", Transaction[i].Product.Interest)
			Transaction[i].TotalTax = (Transaction[i].Amount * Transaction[i].Product.Interest) / 100
			Transaction[i].Total = Transaction[i].TotalTax + Transaction[i].Amount
			filteredTra = append(filteredTra, Transaction[i])
		}
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Transaction retrieved successfully",
		Data:    filteredTra,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
