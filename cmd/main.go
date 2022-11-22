package main

import (
	"github.com/ekholme/go_jwt_prac/srvr"
	"github.com/ekholme/go_jwt_prac/store"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	as := store.NewAuthService()
	us := store.NewUserService()
	ah := srvr.NewAuthHandler(as, us)
	uh := srvr.NewUserHandler(us)

	s := srvr.NewServer(router, ah, uh, as)

	s.Run()

}
