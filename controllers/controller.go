package controller

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/srikanthbhandary/cleanarch/models"
	"github.com/srikanthbhandary/cleanarch/services"
)

type userController struct {
	userService services.UserService
}

type UserController interface {
	GetUsers(c *gin.Context)
	AddUser(c *gin.Context)
}

func (u *userController) GetUsers(c *gin.Context) {
	users, err := u.userService.FindAll()
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
func (u *userController) AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err1 := (u.userService.Validate(&user)); err1 != nil {
		sentry.CaptureException(err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	uid, err := u.userService.Create(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldnot create user in firebase"})
		return
	}
	user.ID = uid.ID
	u.userService.Create(&user)
	c.JSON(http.StatusOK, user)
}

func NewUserController(s services.UserService) UserController {
	return &userController{
		userService: s,
	}
}
