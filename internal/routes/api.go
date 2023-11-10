package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/project1_grup9/internal/adapter/http"
	docs "github.com/iki-rumondor/project1_grup9/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
 	swagFiles "github.com/swaggo/files" 

)

// @title Todo Application
// @version 1.0
// @description This is a todo list management application.
// @host https://project1grup9-production.up.railway.app
// @BasePath /api/v1

func StartTodoServer(handler *customHTTP.TaskHandler) *gin.Engine {
	route := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
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

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swagFiles.Handler))

	return route
}
