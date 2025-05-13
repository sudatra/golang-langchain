package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sudatra/golang-langchain/chains"
);

func generateVacation(r GenerateVacationIdeaRequest) GenerateVacationIdeaResponse {
	id := uuid.New();
	go chains.GenerateVacationIdeaChange(id, r.Budget, r.FavouriteSeason, r.Hobbies);
	return GenerateVacationIdeaResponse{
		Id: id,
		Completed: false,
	}
}

func getVacation(id uuid.UUID) (GetVacationIdeaResponse, error) {
	v, err := chains.GetVacationFromDB(id);
	if err != nil {
		return GetVacationIdeaResponse{}, err;
	}

	return GetVacationIdeaResponse{	
		Id: v.Id,
		Completed: v.Completed,
		Idea: v.Idea,
	}, nil;
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

	registrationRoutes.GET("/:id", func (c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"));
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"});
		} else {
			resp, err := getVacation(id);
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"message": "Id not found"})
			} else {
				c.JSON(http.StatusOK, resp)
			}
		}
	});

	return router;
}