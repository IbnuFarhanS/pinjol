package router

import (
	"net/http"

	"github.com/IbnuFarhanS/pinjol/controller"
	repository "github.com/IbnuFarhanS/pinjol/repository"
	"github.com/gin-gonic/gin"
)

func NewRouter(userRepository repository.UsersRepository, authController *controller.AuthController, usersController *controller.UsersController, rolesController *controller.RolesController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/register", authController.Register)
	authenticationRouter.POST("/login", authController.Login)

	usersRouter := router.Group("/users")
	// usersRouter.GET("", middleware.DeserializeUser(userRepository), usersController.FindAll)
	usersRouter.GET("", usersController.FindAll)
	usersRouter.POST("/", usersController.Insert)

	// rolesRouter := router.Group("/roles")
	// rolesRouter.GET("", rolesController.FindAll)

	return service
}
