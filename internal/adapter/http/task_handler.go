package customHTTP

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

// Method GetAll
func (h *TaskHandler) GetAll(c *gin.Context) {
	tasks, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Method Get ID
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

// Nyoba Delete
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
}
