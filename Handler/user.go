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

		// memecah pesan error
		errors := Helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		respone := Helper.APIRespone("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respone)
		return
	}

	user, err := h.userService.RegisterUser(input)

	if err != nil {
		respone := Helper.APIRespone("Register account failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	formatter := User.FormatUser(user, "token")

	respone := Helper.APIRespone("Accoount has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, respone)
}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukan input (email & Password)
	// input ditangkap handler
	// input di map ke struct LoginUserInput
	// input struct passing service
	// di service mencari dg bantuan repository user dengan email x
	// jika user ditemukan maka cek password

}
