package controller

import (
	"fmt"
	"net/http"

	"github.com/IbnuFarhanS/pinjol/data/request"
	"github.com/IbnuFarhanS/pinjol/data/response"
	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{authService: service}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.ErrorPanic(err)

	token, err_token := controller.authService.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	// ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthController) Register(ctx *gin.Context) {
	createUsersRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	helper.ErrorPanic(err)

	controller.authService.Register(createUsersRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthController) FindAll(ctx *gin.Context) {
	currentUser := ctx.GetString("currentUser")
	users, err := controller.authService.FindAll()
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	filteredUsers := make([]model.Users, 0)
	for _, user := range users {
		if user.Username == currentUser {
			filteredUsers = append(filteredUsers, user)
		}
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Users retrieved successfully",
		Data:    filteredUsers,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
