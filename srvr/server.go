package srvr

import (
	"net/http"

	jwt_prac "github.com/ekholme/go_jwt_prac"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	ah     AuthHandler
	as     jwt_prac.AuthService
	uh     UserHandler
}

// create a new instance of server
// note that this will take some function args
func NewServer(router *gin.Engine, ah AuthHandler, uh UserHandler, as jwt_prac.AuthService) *Server {
	return &Server{
		router: router,
		ah:     ah,
		as:     as,
		uh:     uh,
	}
}

func (s *Server) handleIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "friendo"})
}

// this is not the right place to put this
func (s *Server) Welcome(c *gin.Context) {

	co, err := c.Cookie("demo_cookie")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cookie not found"})
		return
	}

	token, err := s.as.ValidateToken(co)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	claims, ok := token.Claims.(*jwt_prac.CustomClaims)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": claims.Username})
}

func (s *Server) Run() {

	//register routes
	s.router.GET("/", s.handleIndex)
	s.router.POST("/register", s.uh.Register)
	s.router.POST("/login", s.ah.Login)

	//auth endpoints
	authRoutes := s.router.Group("/o", s.ah.ValidateJWT())
	{
		authRoutes.GET("/welcome", s.Welcome)
	}

	s.router.Run(":8080")
}
