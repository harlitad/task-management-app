package handler

import (
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/harlitad/task-management-app/internal/model"
	"github.com/harlitad/task-management-app/internal/usecase"
)

type UserHandler struct {
	UserUsecase usecase.IUserUsecase
}

func NewUserHandler(userUsecase usecase.IUserUsecase) UserHandler {
	return UserHandler{
		UserUsecase: userUsecase,
	}
}

// CreateUser	godoc
// @Summary		Create new user
// @Description	Create new user
// @Tags		auth
// @Accept		json
// @Produce		json
// @Param		user	body		model.CreateUserRequest		true	"Create User JSON"
// @Success		201		{object}	model.CreateUserResponse
// @Router /register [post]
func (h *UserHandler) Create(c *gin.Context) {

	var user model.CreateUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// mapping to user model
	newUser := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	// calling usecase
	res, err := h.UserUsecase.Create(newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model.CreateUserResponse{
		Id:    res.Id,
		Name:  res.Name,
		Email: res.Email,
	})
}

// Authentication	godoc
// @Summary		Authentication
// @Description	Authentication
// @Tags		auth
// @Accept		json
// @Produce		json
// @Param		user	body		model.AuthenticationRequest	true	"User Authentication JSON"
// @Success		200		{object}	model.AuthenticationResponse
// @Router /auth [post]
func (h *UserHandler) Authentication(c *gin.Context) {
	req := model.AuthenticationRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, user, err := h.UserUsecase.Authentication(req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.AuthenticationResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: token,
	})
}
