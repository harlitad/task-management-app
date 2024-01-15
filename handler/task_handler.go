package handler

import (
	"net/http"
	"strconv"

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

func (h *TaskHandler) List(c *gin.Context) {

	result, _ := h.TaskUsecase.List()

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *TaskHandler) Get(c *gin.Context) {

	id := c.Param("id")
	idNumber, _ := strconv.ParseInt(id, 0, 64)

	task, err := h.TaskUsecase.Get(idNumber)
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

func (h *TaskHandler) Create(c *gin.Context) {

	var task model.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := h.TaskUsecase.Create(task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
