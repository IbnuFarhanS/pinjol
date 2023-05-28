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

type UsersController struct {
	usersService service.UsersService
}

func NewUsersController(service service.UsersService) *UsersController {
	return &UsersController{usersService: service}
}

func (c *UsersController) Insert(ctx *gin.Context) {
	createLen := model.Users{}
	err := ctx.ShouldBindJSON(&createLen)
	helper.ErrorPanic(err)

	c.usersService.Save(createLen)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created Users!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UsersController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	updateRol := model.Users{ID: id}
	err = ctx.ShouldBindJSON(&updateRol)
	helper.ErrorPanic(err)

	updatedUsers, err := c.usersService.Update(updateRol)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated Users!",
		Data:    updatedUsers,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *UsersController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	c.usersService.Delete(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted Users!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UsersController) FindAll(ctx *gin.Context) {
	len, err := c.usersService.FindAll()
	helper.ErrorPanic(err)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all Users data!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UsersController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	len, err := c.usersService.FindById(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Users!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UsersController) FindByUsername(ctx *gin.Context) {
	userParam := ctx.Param("username")

	len, err := c.usersService.FindByUsername(userParam)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Users!",
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

