package jwtapp

import (
	httpresponse "app/pkg/http_response"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetToken(ctx *gin.Context) *string {
	if len(strings.Split(ctx.GetHeader("Authorization"), " ")) != 2 {
		httpresponse.Unauthorized(ctx, errors.New("token not found"))
		return nil
	}

	tokenString := strings.Split(ctx.GetHeader("Authorization"), " ")[1]

	return &tokenString
}
