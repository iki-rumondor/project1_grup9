package customHTTP

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project1_grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project1_grup9/internal/application"
)

type TaskHandler struct {
	Service *application.TaskService
}

func NewTaskHandler(service *application.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) GetAll(c *gin.Context) {
	tasks, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	taskID := c.Param("id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Tidak ditemukan"})
		return
	}

	task, err := h.Service.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task Tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Delete(c *gin.Context) {
	taskID := c.Param("id")
	Id, err := strconv.Atoi(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Tidak ditemukan"})
		return
	}

	deletedTask, err := h.Service.Delete(uint(Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus tugas"})
		return
	}

	c.JSON(http.StatusOK, deletedTask)

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

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "task id is not valid",
			"message": err.Error(),
		})
		return
	}

	body.ID = uint(taskId)

	if err := h.Service.UpdateTask(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   fmt.Sprintf("failed to update task with id %d", taskId),
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("task with id %d has been updated successfully", taskId),
	})
}
