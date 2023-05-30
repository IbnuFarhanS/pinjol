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
<<<<<<< HEAD
	"github.com/go-playground/validator/v10"
=======
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)
<<<<<<< HEAD
	validate := validator.New()

	db.Table("users").Find(&model.Users{})
	db.Table("roles").Find(&model.Roles{})
	db.Table("accept_statuses").Find(&model.AcceptStatus{})
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
	authService := service.NewAuthServiceImpl(userRepository)
	usersService := service.NewUsersServiceImpl(userRepository)
	roleService := service.NewRolesServiceImpl(roleRepository)
	accstatService := service.NewAcceptStatusServiceImpl(accstatRepository)
	paymetService := service.NewPaymentMethodServiceImpl(paymetRepository)
	payService := service.NewPaymentsServiceImpl(payRepository)
	proService := service.NewProductsServiceImpl(proRepository)
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
=======

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
	accstatService := service.NewAcceptStatusServiceImpl(accstatRepository, traRepository)
	paymetService := service.NewPaymentMethodServiceImpl(paymetRepository)
	payService := service.NewPaymentServiceImpl(payRepository, traRepository, userRepository, proRepository)
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
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9

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
