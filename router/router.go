package router

import (
	"net/http"

	"github.com/IbnuFarhanS/pinjol/controller"
<<<<<<< HEAD
=======
	"github.com/IbnuFarhanS/pinjol/middleware"
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
func NewRouter(userRepository repository.UsersRepository, authController *controller.AuthController, usersController *controller.UsersController, rolesController *controller.RolesController, accstatController *controller.AcceptStatusController, paymetController *controller.PaymentMethodsController, payController *controller.PaymentsController, proController *controller.ProductsController, traController *controller.TransactionsController) *gin.Engine {
=======
func NewRouter(userRepository repository.UserRepository, authController *controller.AuthController, usersController *controller.UserController, rolesController *controller.RoleController, accstatController *controller.AcceptStatusController, paymetController *controller.PaymentMethodsController, payController *controller.PaymentController, proController *controller.ProductController, traController *controller.TransactionController, uploadKTP *controller.UploadFileKTPController) *gin.Engine {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	r := gin.Default()

	r.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := r.Group("/api")
<<<<<<< HEAD
	// authenticationRouter := router.Group("/auth")
	// authenticationRouter.POST("/register", authController.Register)
	// authenticationRouter.POST("/login", authController.Login)

	usersRouter := router.Group("/users")
	// usersRouter.GET("/", middleware.DeserializeUser(userRepository), authController.FindAll)

	usersRouter.GET("/", usersController.FindAll)
	usersRouter.GET("/:id", usersController.FindByID)
	usersRouter.GET("/username/:username", usersController.FindByUsername)
	usersRouter.PUT("/:id", usersController.Update)
	usersRouter.DELETE("/:id", usersController.Delete)
	usersRouter.POST("/", usersController.Insert)

	rolesRouter := router.Group("/roles")
	rolesRouter.GET("/", rolesController.FindAll)
	rolesRouter.GET("/:id", rolesController.FindByID)
	rolesRouter.GET("/name/:name", rolesController.FindByName)
	rolesRouter.PUT("/:id", rolesController.Update)
	rolesRouter.DELETE("/:id", rolesController.Delete)
	rolesRouter.POST("/", rolesController.Insert)

	accstatRouter := router.Group("/acceptstatus")
	accstatRouter.GET("/", accstatController.FindAll)
	accstatRouter.GET("/:id", accstatController.FindByID)
	accstatRouter.PUT("/:id", accstatController.Update)
	accstatRouter.DELETE("/:id", accstatController.Delete)
	accstatRouter.POST("/", accstatController.Insert)

	payRouter := router.Group("/payments")
	payRouter.GET("/", payController.FindAll)
	payRouter.GET("/:id", payController.FindByID)
	payRouter.PUT("/:id", payController.Update)
	payRouter.DELETE("/:id", payController.Delete)
	payRouter.POST("/", payController.Insert)

	paymetRouter := router.Group("/paymentmethods")
	paymetRouter.GET("/", paymetController.FindAll)
	paymetRouter.GET("/:id", paymetController.FindByID)
	paymetRouter.GET("/name/:name", paymetController.FindByName)
	paymetRouter.PUT("/:id", paymetController.Update)
	paymetRouter.DELETE("/:id", paymetController.Delete)
	paymetRouter.POST("/", paymetController.Insert)

	proRouter := router.Group("/products")
	proRouter.GET("/", proController.FindAll)
	proRouter.GET("/:id", proController.FindByID)
	proRouter.GET("/name/:name", proController.FindByName)
	proRouter.PUT("/:id", proController.Update)
	proRouter.DELETE("/:id", proController.Delete)
	proRouter.POST("/", proController.Insert)

	traRouter := router.Group("/transactions")
	// traRouter.POST("/", middleware.DeserializeUser(userRepository), traController.Insert)
	traRouter.GET("/", traController.FindAll)
	traRouter.GET("/:id", traController.FindByID)
	traRouter.PUT("/:id", traController.Update)
	traRouter.DELETE("/:id", traController.Delete)
	traRouter.POST("/", traController.Insert)
=======
	authenticationRouter := router.Group("/auth")
	{
		authenticationRouter.POST("/register", authController.Register)
		authenticationRouter.POST("/login", authController.Login)
	}

	usersRouter := router.Group("/users")
	usersRouter.GET("/", middleware.DeserializeUser(userRepository), authController.FindAll)

	// usersRouter.GET("/", usersController.FindAll)
	// usersRouter.GET("/:id", usersController.FindByID)
	// usersRouter.PUT("/:id", usersController.Update)
	// usersRouter.DELETE("/:id", usersController.Delete)
	// usersRouter.GET("/username/:username", usersController.FindByUsername)
	// usersRouter.POST("/", usersController.Insert)

	rolesRouter := router.Group("/roles")
	rolesRouter.GET("/", rolesController.FindAll)
	// rolesRouter.GET("/:id", rolesController.FindByID)
	// rolesRouter.GET("/name/:name", rolesController.FindByName)
	// rolesRouter.PUT("/:id", rolesController.Update)
	// rolesRouter.DELETE("/:id", rolesController.Delete)
	rolesRouter.POST("/", rolesController.Insert)

	accstatRouter := router.Group("/acceptstatus")
	accstatRouter.GET("/", middleware.DeserializeUser(userRepository), accstatController.FindAll)
	// accstatRouter.GET("/:id", accstatController.FindByID)
	// accstatRouter.PUT("/:id", accstatController.Update)
	// accstatRouter.DELETE("/:id", accstatController.Delete)
	accstatRouter.POST("/", middleware.DeserializeUser(userRepository), accstatController.Insert)

	payRouter := router.Group("/payments")
	payRouter.POST("/", middleware.DeserializeUser(userRepository), payController.Insert)
	payRouter.GET("/", middleware.DeserializeUser(userRepository), payController.FindAll)
	// payRouter.GET("/:id", payController.FindByID)
	// payRouter.PUT("/:id", payController.Update)
	// payRouter.DELETE("/:id", payController.Delete)
	// payRouter.POST("/", payController.Insert)

	paymetRouter := router.Group("/paymentmethods")
	paymetRouter.GET("/", paymetController.FindAll)
	// paymetRouter.GET("/:id", paymetController.FindByID)
	// paymetRouter.GET("/name/:name", paymetController.FindByName)
	// paymetRouter.PUT("/:id", paymetController.Update)
	// paymetRouter.DELETE("/:id", paymetController.Delete)
	paymetRouter.POST("/", paymetController.Insert)

	proRouter := router.Group("/products")
	// proRouter.POST("/", middleware.DeserializeUser(userRepository), proController.Insert)
	// proRouter.GET("/", proController.FindAll)
	proRouter.GET("/:id", proController.FindByID)
	// proRouter.GET("/name/:name", proController.FindByName)
	// proRouter.PUT("/:id", proController.Update)
	// proRouter.DELETE("/:id", proController.Delete)
	proRouter.POST("/", proController.Insert)

	traRouter := router.Group("/transactions")
	traRouter.POST("/", middleware.DeserializeUser(userRepository), traController.Insert)
	traRouter.GET("/", middleware.DeserializeUser(userRepository), traController.FindAllTransactions)
	traRouter.POST("/export", middleware.DeserializeUser(userRepository), traController.ExportToCSV)
	// traRouter.GET("/", traController.FindAll)
	// traRouter.GET("/:id", traController.FindByID)
	// traRouter.PUT("/:id", traController.Update)
	// traRouter.DELETE("/:id", traController.Delete)
	// traRouter.POST("/", traController.Insert)

	uploadRouter := router.Group("/uploads")
	uploadRouter.POST("/ktp", middleware.DeserializeUser(userRepository), uploadKTP.UploadFileKTP)
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9

	return r
}
