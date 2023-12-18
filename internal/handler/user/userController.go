package user

import (
	"github.com/gin-gonic/gin"

	usecase "project000-backend-user/internal/usecase/user"
)

type (
	Handler interface {
		PostOAuth(c *gin.Context)
	}

	controller struct {
		service usecase.UseCase
	}
)

var _ Handler = (*controller)(nil)

func NewController(service usecase.UseCase) Handler {
	return &controller{service}
}

// PostOAuth
// @Tags /user
// @Router /user/oauth [post]
// @Security BearerAuth
// @Accept json
// @Produce json
//
// @Summary OAuth
//
// @Param PostOauthRequest body string true "oauth authorization code"
//
// @Description OAuth
//
// @Success 200 {object} string
// @Failure 400 {object} string
func (ctrl controller) PostOAuth(c *gin.Context) {
	c.JSON(200, "PostOAuth")
}
