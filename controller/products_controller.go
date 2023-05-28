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

type ProductsController struct {
	productsService service.ProductsService
}

func NewProductsController(service service.ProductsService) *ProductsController {
	return &ProductsController{productsService: service}
}

func (c *ProductsController) Insert(ctx *gin.Context) {
	var product model.Products
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := c.productsService.Save(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created Products!",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *ProductsController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	updatepro := model.Products{ID: id}
	err = ctx.ShouldBindJSON(&updatepro)
	helper.ErrorPanic(err)

	updatedProducts, err := c.productsService.Update(updatepro)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated PaymentMethods!",
		Data:    updatedProducts,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *ProductsController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	c.productsService.Delete(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted PaymentMethods!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *ProductsController) FindAll(ctx *gin.Context) {
	pro, err := c.productsService.FindAll()
	helper.ErrorPanic(err)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all PaymentMethods data!",
		Data:    pro,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *ProductsController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	pro, err := c.productsService.FindById(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched PaymentMethods!",
		Data:    pro,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *ProductsController) FindByName(ctx *gin.Context) {
	userParam := ctx.Param("name")

	pm, err := c.productsService.FindByName(userParam)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched PaymentMethods!",
		Data:    pm,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
