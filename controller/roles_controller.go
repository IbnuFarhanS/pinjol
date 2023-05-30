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

<<<<<<< HEAD
type RolesController struct {
	rolesService service.RolesService
}

func NewRolesController(service service.RolesService) *RolesController {
	return &RolesController{rolesService: service}
}

func (c *RolesController) Insert(ctx *gin.Context) {
	createLen := model.Roles{}
=======
type RoleController struct {
	RoleService service.RoleService
}

func NewRoleController(service service.RoleService) *RoleController {
	return &RoleController{RoleService: service}
}

func (c *RoleController) Insert(ctx *gin.Context) {
	createLen := model.Role{}
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	err := ctx.ShouldBindJSON(&createLen)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

<<<<<<< HEAD
	result, err := c.rolesService.Save(createLen)
=======
	result, err := c.RoleService.Save(createLen)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully created Roles!",
=======
		Message: "Successfully created Role!",
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

<<<<<<< HEAD
func (c *RolesController) Update(ctx *gin.Context) {
=======
func (c *RoleController) Update(ctx *gin.Context) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

<<<<<<< HEAD
	updateRol := model.Roles{ID: id}
	err = ctx.ShouldBindJSON(&updateRol)
	helper.ErrorPanic(err)

	updatedRoles, err := c.rolesService.Update(updateRol)
=======
	updateRol := model.Role{ID: uint(id)}
	err = ctx.ShouldBindJSON(&updateRol)
	helper.ErrorPanic(err)

	updatedRole, err := c.RoleService.Update(updateRol)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully updated Roles!",
		Data:    updatedRoles,
=======
		Message: "Successfully updated Role!",
		Data:    updatedRole,
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	}

	ctx.JSON(http.StatusOK, webResponse)
}
<<<<<<< HEAD
func (c *RolesController) Delete(ctx *gin.Context) {
=======
func (c *RoleController) Delete(ctx *gin.Context) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

<<<<<<< HEAD
	result, err := c.rolesService.Delete(id)
=======
	result, err := c.RoleService.Delete(uint(id))
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully deleted Roles!",
=======
		Message: "Successfully deleted Role!",
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
		Data:    result,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

<<<<<<< HEAD
func (c *RolesController) FindAll(ctx *gin.Context) {
	len, err := c.rolesService.FindAll()
=======
func (c *RoleController) FindAll(ctx *gin.Context) {
	len, err := c.RoleService.FindAll()
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully fetch all Roles data!",
=======
		Message: "Successfully fetch all Role data!",
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

<<<<<<< HEAD
func (c *RolesController) FindByID(ctx *gin.Context) {
=======
func (c *RoleController) FindByID(ctx *gin.Context) {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

<<<<<<< HEAD
	len, err := c.rolesService.FindById(id)
=======
	len, err := c.RoleService.FindById(uint(id))
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully fetched Roles!",
=======
		Message: "Successfully fetched Role!",
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

<<<<<<< HEAD
func (c *RolesController) FindByName(ctx *gin.Context) {
	roleParam := ctx.Param("name")

	len, err := c.rolesService.FindByName(roleParam)
=======
func (c *RoleController) FindByName(ctx *gin.Context) {
	roleParam := ctx.Param("name")

	len, err := c.RoleService.FindByName(roleParam)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
<<<<<<< HEAD
		Message: "Successfully fetched Roles!",
=======
		Message: "Successfully fetched Role!",
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
		Data:    len,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
