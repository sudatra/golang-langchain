package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
);

func generateVacation(r GenerateVacationIdeaRequest) GenerateVacationIdeaResponse {

}

func getVacation(id uuid.UUID) GenerateVacationIdeaResponse {
	
}

func GetVacationRouter(router *gin.Engine) *gin.Engine {
	registrationRoutes := router.Group("/vacation");

	registrationRoutes.POST("/create", func (c *gin.Context) {
		var req GenerateVacationIdeaRequest;
		err := c.BindJSON(&req);
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": "Bad Request"})
		} else {
			c.JSON(http.StatusOK, generateVacation(req));
		}
	});

	registrationRoutes.GET(":/id", func (c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"));
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"});
		} else {
			getVacation(id);
		}
	});
}