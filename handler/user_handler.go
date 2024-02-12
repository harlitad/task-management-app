package handler

import (
	"net/http"

	"strings"

	"example.com/task-management-app/model"
	"example.com/task-management-app/usecase"
	"github.com/gin-gonic/gin"
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
	err := h.UserUsecase.Create(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusCreated, model.CreateUserResponse{
		Name:  user.Name,
		Email: user.Email,
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

	// mapping to user model
	user := model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.UserUsecase.Authentication(user)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.AuthenticationResponse{
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: token,
	})
}
