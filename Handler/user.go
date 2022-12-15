package Handler

import (
	"bwa_startup/Helper"
	"bwa_startup/User"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService User.Service
}

func NewUserHandler(userService User.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct diatas kita passing sebagai parameter service
	var input User.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	user, err := h.userService.RegisterUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	respone := Helper.APIRespone("Accoount has been registered", http.StatusOK, "success", user)

	c.JSON(http.StatusOK, respone)
}
