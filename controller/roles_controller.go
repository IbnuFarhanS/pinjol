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

type RoleController struct {
	RoleService service.RoleService
}

func NewRoleController(service service.RoleService) *RoleController {
	return &RoleController{RoleService: service}
}

func (c *RoleController) Insert(ctx *gin.Context) {
	createLen := model.Role{}
	err := ctx.ShouldBindJSON(&createLen)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.RoleService.Save(createLen)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created Role!",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *RoleController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	updateRol := model.Role{ID: uint(id)}
	err = ctx.ShouldBindJSON(&updateRol)
	helper.ErrorPanic(err)

	updatedRole, err := c.RoleService.Update(updateRol)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated Role!",
		Data:    updatedRole,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *RoleController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	result, err := c.RoleService.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted Role!",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *RoleController) FindAll(ctx *gin.Context) {
	len, err := c.RoleService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all Role data!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *RoleController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	len, err := c.RoleService.FindById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Role!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *RoleController) FindByName(ctx *gin.Context) {
	roleParam := ctx.Param("name")

	len, err := c.RoleService.FindByName(roleParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Role!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
