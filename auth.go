package jwt_prac

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Auth struct {
	Claims *CustomClaims `json:"claims"`
	User   *User         `json:"user"`
	Token  string        `json:"token"`
	Expiry *time.Time    `json:"-"`
}

// i think this is what i want in here
type AuthService interface {
	CreateAuth(u *User) (a *Auth)
	GenerateToken(a *Auth) error
	ValidateToken(tokenStr string) (*jwt.Token, error)
}

//watch https://www.youtube.com/watch?v=p3maH9G_DLM&t=2s for more
