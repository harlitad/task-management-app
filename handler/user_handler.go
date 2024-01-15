package handler

import (
	"net/http"

	"example.com/task-management-app/model"
	"example.com/task-management-app/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.IUserUsecase
}

func (u *UserHandler) Register(c *gin.Context) {

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
