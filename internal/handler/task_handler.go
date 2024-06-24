package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harlitad/task-management-app/internal/dto"
	"github.com/harlitad/task-management-app/internal/model"
	"github.com/harlitad/task-management-app/internal/usecase"
	"github.com/harlitad/task-management-app/pkg/logger"
	"github.com/harlitad/task-management-app/pkg/utils"
)

type TaskHandler struct {
	TaskUsecase usecase.ITaskUsecase
	Logger      logger.ILogger
}

func NewTaskHandler(logger logger.ILogger, taskUsecase usecase.ITaskUsecase) TaskHandler {
	handler := TaskHandler{
		TaskUsecase: taskUsecase,
		Logger:      logger,
	}
	return handler
}

// GetListTask godoc
// @Summary Get list all task
// @Description Get list all task
// @Tags task
// @Accept json
// @Produce json
// @Success 200 {array} dto.Tasks
// @Router /v1/task [get]
// @Security BearerAuth
func (h *TaskHandler) GetTasks(c *gin.Context) {
	h.Logger.Info("incoming request get tasks")
	res, err := h.TaskUsecase.GetTasks(c.Request.Context())
	if err != nil {
		h.Logger.Error("failed get tasks", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseError("failed get tasks", err))
		return
	}
	// process mapping to response
	tasks := make([]dto.Task, 0)
	for _, task := range res {
		data := model.MapToTaskResponse(task)
		tasks = append(tasks, data)
	}

	response := dto.Tasks{
		Data: tasks,
	}
	c.JSON(http.StatusOK, response)
}

// CreateTask godoc
// @Summary Create Task
// @Description Create Task
// @Tags task
// @Accept json
// @Produce json
// @Param		CreateTaskRequest	body		dto.CreateTaskRequest	true	"create task request"
// @Success 201 {array} dto.Task
// @Router /v1/task [post]
// @Security BearerAuth
func (h *TaskHandler) Create(c *gin.Context) {
	// var req dto.CreateTaskRequest

	req := dto.CreateTaskRequest{}

	h.Logger.Info("incoming request create new task")

	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("failed create new task", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseError("failed create new task", err))
		return
	}

	if err := req.Validate(); err != nil {
		h.Logger.Error("failed validate request", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseError("failed validate request", err))
		return
	}

	creatorId, _ := c.Get("userId")

	newTask := model.MapCreateTaskRequest(req)
	newTask.CreatorId = utils.ParseUUID(creatorId.(string))
	createdTask, err := h.TaskUsecase.Create(c.Request.Context(), newTask)
	if err != nil {
		h.Logger.Error("failed create new task", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseError("failed create new task", err))
		return
	}

	response := dto.CreateTaskResponse{
		Task: model.MapToTaskResponse(createdTask),
	}

	c.JSON(http.StatusCreated, response)
}

// GetTaskById godoc
// @Summary Get task by given Id
// @Description Get task by given Id
// @Tags task
// @Accept json
// @Produce json
// @Success 200 {array} dto.Task
// @Param   	id     				path    	string     				true    "task id"
// @Router /v1/task/{id} [get]
// @Security BearerAuth
func (h *TaskHandler) GetTaskById(c *gin.Context) {

	id := c.Param("id")
	h.Logger.WithFields(logger.Fields{"id": id}).Errorf("failed get task by id")

	task, err := h.TaskUsecase.GetTaskById(c.Request.Context(), utils.ParseUUID(id))
	if errors.Is(err, model.ErrorRecordNotFound) {
		h.Logger.WithFields(logger.Fields{"id": id}).Errorf("task not found, %v", err)
		c.AbortWithStatusJSON(http.StatusNotFound, dto.ResponseError("task not found", err))
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseError("failed get task by id", err))
		return
	}

	result := model.MapToTaskResponse(task)

	c.JSON(http.StatusOK, result)
}

// UpdateTask godoc
// @Summary Update task by given Id
// @Description Update task by given Id
// @Tags task
// @Accept json
// @Produce json
// @Param   	id     				path    	string     				true    "task id"
// @Param		UpdateTaskRequest	body		dto.UpdateTaskRequest	true	"update task request"
// @Success 200 {array} dto.Task
// @Router /v1/task/{id} [patch]
// @Security BearerAuth
func (h *TaskHandler) UpdateTask(c *gin.Context) {

	id := c.Param("id")

	req := dto.UpdateTaskRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseError("invalid body request", err))
		return
	}

	if err := req.Validate(); err != nil {
		h.Logger.Error("failed validate request", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseError("failed validate request", err))
		return
	}

	// validate status
	if req.Status != "" {
		reqStatus := model.TaskStatus(req.Status)
		if _, ok := model.TaskStatusMap[reqStatus]; !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseError("failed validate request", errors.New("invalid task status, should be todo, done, inprogress or review")))
			return
		}
	}

	mapUpdatedTask := model.MapUpdateTaskRequest(req)
	mapUpdatedTask.Id = utils.ParseUUID(id)

	task, err := h.TaskUsecase.UpdateTask(c.Request.Context(), mapUpdatedTask)
	if errors.Is(err, model.ErrorRecordNotFound) {
		h.Logger.WithFields(logger.Fields{"id": id}).Errorf("task not found, %v", err)
		c.AbortWithStatusJSON(http.StatusNotFound, dto.ResponseError("task not found", err))
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseError("failed update task by id", err))
		return
	}

	result := model.MapToTaskResponse(task)

	c.JSON(http.StatusOK, result)
}

// DeleteTask godoc
// @Summary Delete task by given id
// @Description Delete task by given id (uuid)
// @Tags task
// @Accept json
// @Produce json
// @Success 204
// @Param   	id     				path    	string     				true    "task id"
// @Router /v1/task/{id} [delete]
// @Security BearerAuth
func (h *TaskHandler) Delete(c *gin.Context) {
	c.Request.Context()

	// context := c.Request.Context()

	id := c.Param("id")

	err := h.TaskUsecase.Delete(c.Request.Context(), utils.ParseUUID(id))
	if errors.Is(err, model.ErrorRecordNotFound) {
		h.Logger.WithFields(logger.Fields{"id": id}).Errorf("task not found, %v", err)
		c.AbortWithStatusJSON(http.StatusNotFound, dto.ResponseError("task not found", err))
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseError("failed delete task by id", err))
		return
	}

	c.Status(http.StatusNoContent)
}

// GetListTaskDueDateToday godoc
// @Summary Get list all task due date by today
// @Description Get list all task due date by today
// @Tags task
// @Accept json
// @Produce json
// @Success 200 {array} dto.Tasks
// @Router /v1/task/today [get]
// @Security BearerAuth
func (h *TaskHandler) GetTasksToday(c *gin.Context) {
	h.Logger.Info("incoming request get tasks")
	res, err := h.TaskUsecase.GetTasksToday(c.Request.Context())
	if err != nil {
		h.Logger.Error("failed get tasks", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseError("failed get tasks", err))
		return
	}
	// process mapping to response
	tasks := make([]dto.Task, 0)
	for _, task := range res {
		data := model.MapToTaskResponse(task)
		tasks = append(tasks, data)
	}

	response := dto.Tasks{
		Data: tasks,
	}
	c.JSON(http.StatusOK, response)
}
