package store

import (
	"errors"
	"time"

	jwt_prac "github.com/ekholme/go_jwt_prac"
	"github.com/golang-jwt/jwt/v4"
)

const (
	secretKey = "vryscrtkey"
)

type authService struct {
	key string
}

func NewAuthService() jwt_prac.AuthService {
	return &authService{
		key: secretKey,
	}
}

// methods
func (as *authService) CreateAuth(u *jwt_prac.User) (a *jwt_prac.Auth) {

	exp := time.Now().Add(2 * time.Hour).Unix()

	claims := &jwt_prac.CustomClaims{
		Username: u.Username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "sleazy_e",
			ExpiresAt: exp,
		},
	}

	return &jwt_prac.Auth{
		User:   u,
		Claims: claims,
	}
}

// this should write out the token to the auth struct
// without having to explicitly return a string
func (as *authService) GenerateToken(a *jwt_prac.Auth) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.Claims)

	tokenStr, err := token.SignedString([]byte(as.key))

	if err != nil {
		return err
	}

	a.Token = tokenStr

	return nil
}

func (as *authService) ValidateToken(tokenStr string) (*jwt.Token, error) {

	tkn, err := jwt.ParseWithClaims(tokenStr, &jwt_prac.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(as.key), nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, errors.New("token not valid")
	}

	return tkn, nil
}
