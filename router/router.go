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
	r.DB, _ = db.NewDB()
	err := r.DB.MigrationDB()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	h := handler.NewHTTPHandler(r.DB)
	userWeb := r.Engine.Group("/user")
	{
		userWeb.GET("/:id", h.GetUserProfile)
		userWeb.POST("/create", h.CreateNewUser)
	}
	return nil
}
func NewRouter() Router {
	var r Router
	r.Setup()

	return r
}
