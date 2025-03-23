package httpresponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Unauthorized(ctx *gin.Context, err error) {
	res := Response{
		Data:    nil,
		Message: err.Error(),
		Status:  http.StatusUnauthorized,
		Error:   err,
	}

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
}

func BadRequest(ctx *gin.Context, err error) {
	res := Response{
		Data:    nil,
		Message: err.Error(),
		Status:  http.StatusBadRequest,
		Error:   err,
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func Success(ctx *gin.Context, data interface{}) {
	res := Response{
		Data:    data,
		Message: "",
		Status:  http.StatusOK,
		Error:   nil,
	}

	ctx.JSON(http.StatusOK, res)
}

func InternalServerError(ctx *gin.Context, err error) {
	res := Response{
		Data:    nil,
		Message: err.Error(),
		Status:  http.StatusInternalServerError,
		Error:   err,
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
}
