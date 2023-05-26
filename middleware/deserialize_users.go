package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/IbnuFarhanS/pinjol/config"
	"github.com/IbnuFarhanS/pinjol/helper"
	reposity "github.com/IbnuFarhanS/pinjol/repository"
	"github.com/IbnuFarhanS/pinjol/utils"
	"github.com/gin-gonic/gin"
)

func DeserializeUser(userRepository reposity.UsersRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		id, err_id := strconv.ParseInt(fmt.Sprint(sub), 10, 64)
		helper.ErrorPanic(err_id)
		result, err := userRepository.FindById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", result.Username)
		ctx.Next()

	}
}
