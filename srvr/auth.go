package srvr

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	jwt_prac "github.com/ekholme/go_jwt_prac"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context)
	ValidateJWT() gin.HandlerFunc
}

type authHandler struct {
	authService jwt_prac.AuthService
	userService jwt_prac.UserService
}

func NewAuthHandler(as jwt_prac.AuthService, us jwt_prac.UserService) AuthHandler {
	return &authHandler{
		authService: as,
		userService: us,
	}
}

// handler to perform login
func (ah authHandler) Login(c *gin.Context) {
	var user *jwt_prac.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//i don't really like that I have to do this, but w/e
	//users := ah.userService.GetAllUsers()

	users := []*jwt_prac.User{
		&jwt_prac.User{
			Username: "erice",
			Password: "pass123",
		},
	}

	ind, err := CheckUserExists(user, users)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ref := users[ind]

	err = CheckPwMatch(user, ref)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a := ah.authService.CreateAuth(user)

	err = ah.authService.GenerateToken(a)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	exp := int(a.Claims.ExpiresAt - time.Now().Unix())

	c.SetCookie("demo_cookie", a.Token, exp, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "logged in successfully!"})
}

// this is the auth middleware
func (ah authHandler) ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		co, err := c.Cookie("demo_cookie")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cookie not found"})
			return
		}

		token, err := ah.authService.ValidateToken(co)

		if token.Valid {
			claims := token.Claims.(*jwt_prac.CustomClaims)
			fmt.Println("Claims:", claims)
		} else {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "you can't be here"})
		}
	}
}

// helper funcs
// checking that a user exists in the slice of users
func CheckUserExists(u *jwt_prac.User, s []*jwt_prac.User) (int, error) {
	for k, v := range s {
		if u.Username == v.Username {
			return k, nil
		}
	}
	return 0, errors.New("user doesn't exist")
}

// check that passwords match
func CheckPwMatch(inp *jwt_prac.User, ref *jwt_prac.User) error {

	a := (inp.Password)
	b := (ref.Password)

	//err := bcrypt.CompareHashAndPassword(a, b)

	if a != b {
		return errors.New("passwords don't match")
	} else {
		return nil
	}

}
