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
	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").Find(&model.Users{})
	db.Table("roles").Find(&model.Roles{})
	db.Table("accept_status").Find(&model.AcceptStatus{})
	db.Table("products").Find(&model.Products{})
	db.Table("payments").Find(&model.Payments{})
	db.Table("payment_methods").Find(&model.PaymentMethod{})
	db.Table("transactions").Find(&model.Transactions{})

	//Init Repository
	userRepository := repository.NewUsersRepositoryImpl(db)
	roleRepository := repository.NewRolesRepositoryImpl(db)
	accstatRepository := repository.NewAcceptStatusRepositoryImpl(db)
	paymetRepository := repository.NewPaymentMethodRepositoryImpl(db)
	payRepository := repository.NewPaymentsRepositoryImpl(db)
	proRepository := repository.NewProductsRepositoryImpl(db)
	traRepository := repository.NewTransactionsRepositoryImpl(db)

	//Init Service
	authService := service.NewAuthServiceImpl(userRepository, validate)
	usersService := service.NewUsersServiceImpl(userRepository, validate)
	roleService := service.NewRolesServiceImpl(roleRepository, validate)
	accstatService := service.NewAcceptStatusServiceImpl(accstatRepository, validate)
	paymetService := service.NewPaymentMethodServiceImpl(paymetRepository, validate)
	payService := service.NewPaymentsServiceImpl(payRepository, validate)
	proService := service.NewProductsServiceImpl(proRepository, validate)
	traService := service.NewTransactionsServiceImpl(traRepository, validate)

	//Init controller
	authController := controller.NewAuthController(authService)
	usersController := controller.NewUsersController(usersService)
	rolesController := controller.NewRolesController(roleService)
	accstatController := controller.NewAcceptStatusController(accstatService)
	paymetController := controller.NewPaymentMethodsController(paymetService)
	payController := controller.NewPaymentsController(payService)
	proController := controller.NewProductsController(proService)
	traController := controller.NewTransactionsController(traService)

	routes := router.NewRouter(userRepository, authController, usersController, rolesController, accstatController, paymetController, payController, proController, traController)

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
