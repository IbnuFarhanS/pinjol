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

type RolesController struct {
	rolesService service.RolesService
}

func NewRolesController(service service.RolesService) *RolesController {
	return &RolesController{rolesService: service}
}

func (c *RolesController) Insert(ctx *gin.Context) {
	createLen := model.Roles{}
	err := ctx.ShouldBindJSON(&createLen)
	helper.ErrorPanic(err)

	c.rolesService.Save(createLen)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created Roles!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *RolesController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	updateRol := model.Roles{ID: id}
	err = ctx.ShouldBindJSON(&updateRol)
	helper.ErrorPanic(err)

	updatedRoles, err := c.rolesService.Update(updateRol)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated Roles!",
		Data:    updatedRoles,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *RolesController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	c.rolesService.Delete(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted Roles!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *RolesController) FindAll(ctx *gin.Context) {
	len, err := c.rolesService.FindAll()
	helper.ErrorPanic(err)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all Roles data!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *RolesController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	len, err := c.rolesService.FindById(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Roles!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *RolesController) FindByName(ctx *gin.Context) {
	roleParam := ctx.Param("name")

	len, err := c.rolesService.FindByName(roleParam)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Roles!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
