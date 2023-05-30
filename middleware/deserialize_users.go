package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/IbnuFarhanS/pinjol/config"
	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/IbnuFarhanS/pinjol/utils"
	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
func DeserializeUser(userRepository repository.UsersRepository) gin.HandlerFunc {
=======
func DeserializeUser(userRepository repository.UserRepository) gin.HandlerFunc {
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")

		fields := strings.Fields(authorizationHeader)

		if len(fields) != 2 || fields[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Unauthorization"})
			return
		}

		token = fields[1]

		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		id, err_id := strconv.ParseInt(fmt.Sprint(sub), 10, 64)
		helper.ErrorPanic(err_id)
<<<<<<< HEAD
		result, err := userRepository.FindById(id)
=======
		result, err := userRepository.FindById(uint(id))
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", result.Username)
		ctx.Set("currentUserID", id)
		ctx.Next()
	}
}
