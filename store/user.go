package store

import (
	jwt_prac "github.com/ekholme/go_jwt_prac"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	Users []*jwt_prac.User
}

// initiate service
func NewUserService() jwt_prac.UserService {
	return &userService{
		Users: nil,
	}
}

// method to create a user
func (us *userService) CreateUser(u *jwt_prac.User) []*jwt_prac.User {
	res := append(us.Users, u)

	return res
}

func (us *userService) GetAllUsers() []*jwt_prac.User {
	return us.Users
}

// helper function to hash a password on registration
func HashPw(u *jwt_prac.User) error {
	hp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hp)

	return nil
}
