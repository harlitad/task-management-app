package handler

import (
	"fmt"
	"net/http"

	"example.com/task-management-app/model"
	"example.com/task-management-app/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	TaskUsecase usecase.ITaskUsecase
}

func NewTaskHandler(u usecase.ITaskUsecase) TaskHandler {
	handler := TaskHandler{
		TaskUsecase: u,
	}
	return handler
}

// ListTask godoc
// @Summary Get list user's task
// @Description Get list user's task
// @Tags task
// @Accept json
// @Produce json
// @Success 200 {array} model.Task
// @Router /v1/task [get]
// @Security BearerAuth
func (h *TaskHandler) List(c *gin.Context) {
	userId, _ := c.Get("userId")

	result, _ := h.TaskUsecase.List(userId.(string))

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// GetTask godoc
// @Summary Get user's task by id
// @Description Get user's task by id
// @Tags task
// @Accept json
// @Produce json
// @Success 200 {object} model.Task
// @Router /v1/task/{id} [get]
// @Param	id	path	string	true	"Task ID"
// @Security BearerAuth
func (h *TaskHandler) Get(c *gin.Context) {

	id := c.Param("id")
	userId, _ := c.Get("userId")

	task, err := h.TaskUsecase.Get(id, userId.(string))
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "record not found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, task)
}

// CreateTask godoc
// @Summary Create user's task
// @Description Create user's task
// @Tags task
// @Accept json
// @Produce json
// @Param		user	body		model.CreateTaskRequest	true	"Create Task"
// @Success		201		{object}	model.CreateTaskResponse
// @Router /v1/task [post]
// @Security BearerAuth
func (h *TaskHandler) Create(c *gin.Context) {

	var task model.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId, _ := c.Get("userId")
	task.UserId = userId.(string)

	res, err := h.TaskUsecase.Create(task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	fmt.Println(res.Id)

	c.JSON(http.StatusCreated, res)
}

// DeleteTask godoc
// @Summary Delete user's task by id
// @Description Delete user's task by id
// @Tags task
// @Accept json
// @Produce json
// @Success 204
// @Router /v1/task/{id} [delete]
// @Param	id	path	string	true	"Task ID"
// @Security BearerAuth
func (h *TaskHandler) Delete(c *gin.Context) {

	id := c.Param("id")
	userId, _ := c.Get("userId")

	err := h.TaskUsecase.Delete(id, userId.(string))
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "record not found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	c.Status(http.StatusNoContent)
}
