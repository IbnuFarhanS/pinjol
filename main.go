package main

import (
	"log"
	"net/http"
	"time"

	"github.com/IbnuFarhanS/pinjol/config"
	"github.com/IbnuFarhanS/pinjol/controller"
	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/IbnuFarhanS/pinjol/router"
	"github.com/IbnuFarhanS/pinjol/service"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Role{})
	db.AutoMigrate(&model.AcceptStatus{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Payment{})
	db.AutoMigrate(&model.PaymentMethod{})
	db.AutoMigrate(&model.Transaction{})

	//Init Repository
	userRepository := repository.NewUserRepositoryImpl(db)
	roleRepository := repository.NewRoleRepositoryImpl(db)
	accstatRepository := repository.NewAcceptStatusRepositoryImpl(db)
	paymetRepository := repository.NewPaymentMethodRepositoryImpl(db)
	payRepository := repository.NewPaymentRepositoryImpl(db)
	proRepository := repository.NewProductRepositoryImpl(db)
	traRepository := repository.NewTransactionRepositoryImpl(db)

	//Init Service
	authService := service.NewAuthServiceImpl(userRepository)
	usersService := service.NewUserServiceImpl(userRepository)
	roleService := service.NewRoleServiceImpl(roleRepository)
	accstatService := service.NewAcceptStatusServiceImpl(accstatRepository)
	paymetService := service.NewPaymentMethodServiceImpl(paymetRepository)
	payService := service.NewPaymentServiceImpl(payRepository)
	proService := service.NewProductServiceImpl(proRepository)
	traService := service.NewTransactionServiceImpl(traRepository, userRepository)
	uploadService := service.NewUploadFileKTPService()

	//Init controller
	authController := controller.NewAuthController(authService)
	usersController := controller.NewUserController(usersService)
	rolesController := controller.NewRoleController(roleService)
	accstatController := controller.NewAcceptStatusController(accstatService)
	paymetController := controller.NewPaymentMethodsController(paymetService)
	payController := controller.NewPaymentController(payService)
	proController := controller.NewProductController(proService)
	traController := controller.NewTransactionController(traService)
	uploadController := controller.NewUploadFileKTPController(uploadService)

	routes := router.NewRouter(userRepository, authController, usersController, rolesController, accstatController, paymetController, payController, proController, traController, uploadController)

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
