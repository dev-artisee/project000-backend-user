package user

import (
	"github.com/gin-gonic/gin"

	"project000-backend-user/config"
)

func Route(group *gin.RouterGroup, handler Handler, config config.Config) {
	group.POST("/oauth", handler.PostOAuth)
}
