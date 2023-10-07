package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/project1_grup9/internal/adapter/http"
)

func StartTodoServer(handler *customHTTP.TaskHandler) *gin.Engine{
	route := gin.Default()

	route.Group("/api/v1")
	route.GET("/todos")
	route.POST("/todos")
	route.GET("/todos/:id")
	route.PUT("/todos/:id")
	route.DELETE("/todos/:id")

	return route
}