package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/project1_grup9/internal/adapter/http"
)

func StartTodoServer(handler *customHTTP.TaskHandler) *gin.Engine {
	route := gin.Default()

	v1 := route.Group("/api/v1")
	{
		v1.GET("/todos", handler.GetAll)
		v1.POST("/todos")
		v1.GET("/todos/:id", handler.GetTaskByID)
		v1.PUT("/todos/:id")
		v1.DELETE("/todos/:id")
	}
	return route
}
