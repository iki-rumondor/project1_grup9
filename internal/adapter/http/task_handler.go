package customHTTP

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project1_grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project1_grup9/internal/application"
	"gorm.io/gorm"
)

type TaskHandler struct {
	Service *application.TaskService
}

func NewTaskHandler(service *application.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to get all tasks",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": tasks,
	})
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "task id is not valid",
			"message": err.Error(),
		})
		return
	}

	task, err := h.Service.GetByTaskID(uint(taskID))
	if errors.Is(gorm.ErrRecordNotFound, err) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("task with id %d is not found", taskID),
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   fmt.Sprintf("failed to get task with id %d", taskID),
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": task,
	})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "task id is not valid",
			"message": err.Error(),
		})
		return
	}

	if err := h.Service.DeleteTask(uint(taskID)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   fmt.Sprintf("failed to update task with id %d", taskID),
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("task with id %d has been deleted successfully", taskID),
	})
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var body request.CreateTask

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "failed to parse request body",
			"message": err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "request body is not valid",
			"message": err.Error(),
		})
		return
	}

	if err := h.Service.CreateTask(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to create new task",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "new task has been created successfully",
	})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var body request.UpdateTask

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "failed to parse request body",
			"message": err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "request body is not valid",
			"message": err.Error(),
		})
		return
	}

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "task id is not valid",
			"message": err.Error(),
		})
		return
	}

	body.ID = uint(taskID)

	if err := h.Service.UpdateTask(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   fmt.Sprintf("failed to update task with id %d", taskID),
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("task with id %d has been updated successfully", taskID),
	})
}
