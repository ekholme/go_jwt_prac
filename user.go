package jwt_prac

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserService interface {
	CreateUser(u *User)
}
