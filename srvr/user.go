package srvr

import (
	"net/http"

	jwt_prac "github.com/ekholme/go_jwt_prac"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
}

type userHandler struct {
	userService jwt_prac.UserService
}

// create new instance of userHandler
func NewUserHandler(us jwt_prac.UserService) UserHandler {
	return &userHandler{
		userService: us,
	}
}

// register handler
func (uh userHandler) Register(c *gin.Context) {
	var user jwt_prac.User

	//var users []*jwt_prac.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//err = store.HashPw(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_ = uh.userService.CreateUser(&user)

	//res := append(users, &user)

	c.JSON(http.StatusOK, gin.H{
		"message":   "user created",
		"user list": uh.userService.GetAllUsers(),
	})

}
