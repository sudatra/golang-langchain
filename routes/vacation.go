package routes

import (
	"github.com/gin-gonic/gin"
)

func GetVacationRouter(router *gin.Engine) *gin.Engine {
	registrationRoutes := router.Group("/vacation");

	registrationRoutes.POST("/create", func (c *gin.Context) {
		var request 
	})
}