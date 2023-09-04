package routes

import (
	"e-wallet/src/handler"

	"github.com/gin-gonic/gin"
)

func (r *Router) Auth(route *gin.RouterGroup, h *handler.Handler){
	route.POST("/register", h.)
}