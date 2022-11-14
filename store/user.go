package store

import (
	"errors"

	jwt_prac "github.com/ekholme/go_jwt_prac"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

// initiate service
func NewUserService() jwt_prac.UserService {
	return &userService{}
}

// method to create a user
func (user *userService) CreateUser(u *jwt_prac.User, s []*jwt_prac.User) []*jwt_prac.User {
	res := append(s, u)

	return res
}

//some helper functions

// first is a pw hash func
func HashPw(u *jwt_prac.User) error {
	hp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hp)

	return nil
}

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
func CheckPwMatch(inp *jwt_prac.User, ex *jwt_prac.User) error {
	//TODO
}
