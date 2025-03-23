package middlewareapp

import (
	constant "app/internal/constants"
	httpresponse "app/pkg/http_response"
	jwtapp "app/pkg/jwt"
	logapp "app/pkg/log"

	"github.com/gin-gonic/gin"
)

func GetProfileId(ctx *gin.Context) {
	profileId, err := jwtapp.GetProfileId(ctx)
	if err != nil {
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("get-profile-id", err.Error(), constant.ERROR_LOG)
		return
	}
	ctx.Set("profile_id", profileId)

	ctx.Next()
}
