package controller

import (
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{ProductService: service}
}

func (c *ProductController) Insert(ctx *gin.Context) {
	createpro := model.Product{}
	err := ctx.ShouldBindJSON(&createpro)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.ProductService.Save(createpro)

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

func (c *ProductController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatepro := model.Product{ID: uint(id)}
	err = ctx.ShouldBindJSON(&updatepro)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := c.ProductService.Update(updatepro)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated PaymentMethods!",
		Data:    updatedProduct,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *ProductController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.ProductService.Delete(uint(id))
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

func (c *ProductController) FindAll(ctx *gin.Context) {
	pro, err := c.ProductService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all PaymentMethods data!",
		Data:    pro,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *ProductController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pro, err := c.ProductService.FindById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched PaymentMethods!",
		Data:    pro,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *ProductController) FindByName(ctx *gin.Context) {
	proParam := ctx.Param("name")

	pro, err := c.ProductService.FindByName(proParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched PaymentMethods!",
		Data:    pro,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
