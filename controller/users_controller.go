package controller

import (
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
type UsersController struct {
	usersService service.UsersService
}

func NewUsersController(service service.UsersService) *UsersController {
	return &UsersController{usersService: service}
}

func (c *UsersController) Insert(ctx *gin.Context) {
	createLen := model.Users{}
	err := ctx.ShouldBindJSON(&createLen)
=======
type UserController struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{UserService: service}
}

func (c *UserController) Insert(ctx *gin.Context) {
	createuser := model.User{}
	err := ctx.ShouldBindJSON(&createuser)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

<<<<<<< HEAD
	result, err := c.usersService.Save(createLen)
=======
	result, err := c.UserService.Save(createuser)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully created Users!",
=======
		Message: "Successfully created User!",
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

<<<<<<< HEAD
func (c *UsersController) Update(ctx *gin.Context) {
=======
func (c *UserController) Update(ctx *gin.Context) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

<<<<<<< HEAD
	updateRol := model.Users{ID: id}
=======
	updateRol := model.User{ID: uint(id)}
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	err = ctx.ShouldBindJSON(&updateRol)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

<<<<<<< HEAD
	updatedUsers, err := c.usersService.Update(updateRol)
=======
	updatedUser, err := c.UserService.Update(updateRol)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully updated Users!",
		Data:    updatedUsers,
=======
		Message: "Successfully updated User!",
		Data:    updatedUser,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}

	ctx.JSON(http.StatusOK, webResponse)
}
<<<<<<< HEAD
func (c *UsersController) Delete(ctx *gin.Context) {
=======
func (c *UserController) Delete(ctx *gin.Context) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

<<<<<<< HEAD
	c.usersService.Delete(id)
=======
	del,err := c.UserService.Delete(uint(id))
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully deleted Users!",
		Data:    nil,
=======
		Message: "Successfully deleted User!",
		Data:    del,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}

	ctx.JSON(http.StatusOK, webResponse)
}

<<<<<<< HEAD
func (c *UsersController) FindAll(ctx *gin.Context) {
	len, err := c.usersService.FindAll()
=======
func (c *UserController) FindAll(ctx *gin.Context) {
	user, err := c.UserService.FindAll()
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully fetch all Users data!",
		Data:    len,
=======
		Message: "Successfully fetch all User data!",
		Data:    user,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}

	ctx.JSON(http.StatusOK, webResponse)
}

<<<<<<< HEAD
func (c *UsersController) FindByID(ctx *gin.Context) {
=======
func (c *UserController) FindByID(ctx *gin.Context) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

<<<<<<< HEAD
	len, err := c.usersService.FindById(id)
=======
	user, err := c.UserService.FindById(uint(id))
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully fetched Users!",
		Data:    len,
=======
		Message: "Successfully fetched User!",
		Data:    user,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}

	ctx.JSON(http.StatusOK, webResponse)
}

<<<<<<< HEAD
func (c *UsersController) FindByUsername(ctx *gin.Context) {
	userParam := ctx.Param("username")

	len, err := c.usersService.FindByUsername(userParam)
=======
func (c *UserController) FindByUsername(ctx *gin.Context) {
	userParam := ctx.Param("username")

	user, err := c.UserService.FindByUsername(userParam)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully fetched Users!",
		Data:    len,
=======
		Message: "Successfully fetched User!",
		Data:    user,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}

	ctx.JSON(http.StatusOK, webResponse)
}
