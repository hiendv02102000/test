package router

import (
	"fmt"
	"test/db"
	"test/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	DB     db.Database
}

func (r *Router) Setup() error {
	r.Engine = gin.Default()
	DB, err := db.NewDB()
	r.DB = DB
	if err != nil {
		fmt.Print(err)
		return err
	}
	h := handler.NewHTTPHandler(r.DB)
	userWeb := r.Engine.Group("/user")
	{
		userWeb.GET("/:id", h.GetUserProfile)
		userWeb.POST("/register", h.RegisterUser)
	}
	return nil
}
func NewRouter() Router {
	var r Router
	r.Setup()

	return r
}
