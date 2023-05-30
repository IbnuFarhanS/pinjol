package controller

import (
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{UserService: service}
}

func (c *UserController) Insert(ctx *gin.Context) {
	createuser := model.User{}
	err := ctx.ShouldBindJSON(&createuser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := c.UserService.Save(createuser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created User!",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UserController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updateRol := model.User{ID: uint(id)}
	err = ctx.ShouldBindJSON(&updateRol)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatedUser, err := c.UserService.Update(updateRol)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated User!",
		Data:    updatedUser,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *UserController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	del, err := c.UserService.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted User!",
		Data:    del,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UserController) FindAll(ctx *gin.Context) {
	user, err := c.UserService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all User data!",
		Data:    user,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UserController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := c.UserService.FindById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched User!",
		Data:    user,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UserController) FindByUsername(ctx *gin.Context) {
	userParam := ctx.Param("username")

	user, err := c.UserService.FindByUsername(userParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched User!",
		Data:    user,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
