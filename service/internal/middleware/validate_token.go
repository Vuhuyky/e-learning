package middlewareapp

import (
	httpresponse "app/pkg/http_response"
	jwtapp "app/pkg/jwt"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func ValidateToken(ctx *gin.Context) {
	if len(strings.Split(ctx.GetHeader("Authorization"), " ")) != 2 {
		httpresponse.Unauthorized(ctx, errors.New("token not found"))
		return
	}

	tokenString := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	mapData, errMapData := jwtapp.GetClaim(tokenString)

	if errMapData != nil {
		httpresponse.Unauthorized(ctx, errMapData)
		return
	}

	exp := mapData["exp"].(float64)

	if time.Now().Unix() > int64(exp) {
		httpresponse.Unauthorized(ctx, errors.New("token expired"))
		return
	}

	ctx.Next()
}
