package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sudatra/golang-langchain/routes"
)

func main() {
	r := gin.Default();

	routes.GetVacationRouter(r);
	r.Run();
}