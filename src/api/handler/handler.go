package handler

import (
	"go-clean-arch/src/domain"
	helpers "go-clean-arch/src/helper"
	"go-clean-arch/src/middleware"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	Entity domain.Entity
}

func InitHandler(r *gin.RouterGroup, e domain.Entity) {
	handler := &AppHandler{
		Entity: e,
	}

	r.GET("/", handler.Home)

	users := r.Group("/user")
	{
		users.POST("/login", handler.Login)
		users.GET("/info", middleware.Auth(handler.Entity), handler.GetUserByUuid)
	}
}

func (a *AppHandler) Home(c *gin.Context) {
	params := map[string]interface{}{
		"payload": gin.H{"message": "OK", "version": "2"},
		"meta":    gin.H{"message": "OK"},
	}
	c.JSON(200, helpers.OutputAPIResponseWithPayload(params))
}
