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

// TaskHandler handles HTTP requests related to tasks.
type TaskHandler struct {
	Service *application.TaskService
}

type responseMessage struct {
	Message string `json:"message"`
}

func NewTaskHandler(service *application.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

// GetAllTasks godoc
// @Summary Get all tasks
// @Description Retrieve a list of all tasks.
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {object} []response.Task
// @Failure 500 {object} responseMessage
// @Router /todos [get]
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseMessage{
			Message: "failed to get all tasks",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": tasks,
	})
}

// GetTaskByID godoc
// @Summary Get a task by ID
// @Description Retrieve a task by its unique identifier.
// @Tags todos
// @Accept json
// @Produce json
// @Param id path uint true "Task ID"
// @Success 200 {object} response.Task
// @Failure 500 {object} responseMessage
// @Failure 404 {object} responseMessage
// @Router /todos/{id} [get]
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseMessage{
			Message: "task id is not valid",
		})
		return
	}

	task, err := h.Service.GetByTaskID(uint(taskID))
	if errors.Is(gorm.ErrRecordNotFound, err) {
		c.AbortWithStatusJSON(http.StatusNotFound, responseMessage{
			Message: fmt.Sprintf("task with id %d is not found", taskID),
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseMessage{
			Message: fmt.Sprintf("failed to get task with id %d", taskID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": task,
	})
}

// DeleteTask godoc
// @Summary Delete a task by ID
// @Description Delete a task by its unique identifier.
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} responseMessage
// @Failure 400 {object} responseMessage
// @Failure 500 {object} responseMessage
// @Router /todos/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseMessage{
			Message: "task id is not valid",
		})
		return
	}

	err = h.Service.DeleteTask(uint(taskID))

	if errors.Is(gorm.ErrRecordNotFound, err) {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseMessage{
			Message: fmt.Sprintf("task with id %d is not found", taskID),
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseMessage{
			Message: fmt.Sprintf("something wrong when delete task with id %d", taskID),
		})
		return
	}

	c.JSON(http.StatusOK, responseMessage{
		Message: fmt.Sprintf("task with id %d has been deleted successfully", taskID),
	})
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with the provided data.
// @Tags todos
// @Accept json
// @Produce json
// @Param request.Task body request.Task true "create task"
// @Success 201 {object} responseMessage
// @Failure 400 {object} responseMessage
// @Failure 500 {object} responseMessage
// @Router /todos [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var body request.Task

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseMessage{
			Message: "failed to parse request body",
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseMessage{
			Message: "request body is not valid",
		})
		return
	}

	if err := h.Service.CreateTask(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseMessage{
			Message: "failed to create new task",
		})
		return
	}

	c.JSON(http.StatusCreated, responseMessage{
		Message: "new task has been created successfully",
	})
}

// UpdateTask godoc
// @Summary Update an existing task
// @Description Update an existing task with the provided data.
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param request.Task body request.Task true "Updated Task Data"
// @Success 200 {object} responseMessage
// @Failure 400 {object} responseMessage
// @Failure 500 {object} responseMessage
// @Router /todos/{id} [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var body request.TaskWithID

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseMessage{
			Message: "failed to parse request body",
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseMessage{
			Message: "request body is not valid",
		})
		return
	}

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseMessage{
			Message: "task id is not valid",
		})
		return
	}

	body.ID = uint(taskID)

	err = h.Service.UpdateTask(&body)

	if errors.Is(gorm.ErrRecordNotFound, err) {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseMessage{
			Message: fmt.Sprintf("task with id %d is not found", taskID),
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseMessage{
			Message: fmt.Sprintf("something wrong when update task with id %d", taskID),
		})
		return
	}

	c.JSON(http.StatusOK, responseMessage{
		Message: fmt.Sprintf("task with id %d has been updated successfully", taskID),
	})
}
