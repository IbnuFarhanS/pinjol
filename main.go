package main

import (
	"log"
	"net/http"
	"time"

	"github.com/IbnuFarhanS/pinjol/config"
	"github.com/IbnuFarhanS/pinjol/controller"
	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	reposity "github.com/IbnuFarhanS/pinjol/repository"
	"github.com/IbnuFarhanS/pinjol/router"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").Find(&model.Users{})
	db.Table("roles").Find(&model.Roles{})

	//Init Repository
	userRepository := reposity.NewUsersRepositoryImpl(db)
	roleRepository := reposity.NewRolesRepositoryImpl(db)

	//Init Service
	authService := service.NewAuthServiceImpl(userRepository, validate)
	usersService := service.NewUsersServiceImpl(userRepository, validate)
	roleService := service.NewRolesServiceImpl(roleRepository, validate)

	//Init controller
	authController := controller.NewAuthController(authService)
	usersController := controller.NewUsersController(usersService)
	rolesController := controller.NewRolesController(roleService)

	routes := router.NewRouter(userRepository, authController, usersController, rolesController)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
