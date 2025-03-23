package routerconfig

import (
	constant "app/internal/constants"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	Endpoint   string
	Method     constant.HTTP_METHOD
	Middleware []gin.HandlerFunc
	Handle     gin.HandlerFunc
}
