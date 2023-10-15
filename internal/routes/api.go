package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/project1_grup9/internal/adapter/http"
)

// @title API
// @version 1
// @Summary Start the Todo API server
// @Description Initialize the server and set up routes for the Todo API.

// @host localhost:8080
// @BasePath /api/v1
func StartTodoServer(handler *customHTTP.TaskHandler) *gin.Engine {
	route := gin.Default()

	v1 := route.Group("/api/v1")
	{
		// FindAll
		v1.GET("/todos", handler.GetAllTasks)

		// Create
		v1.POST("/todos", handler.CreateTask)

		// FindByID
		v1.GET("/todos/:id", handler.GetTaskByID)

		// Update
		v1.PUT("/todos/:id", handler.UpdateTask)

		// Delete
		v1.DELETE("/todos/:id", handler.DeleteTask)
	}

	return route
}
