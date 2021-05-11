package router

import (
	"test/db"
	"test/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	DB     db.Database
}

func (r *Router) Setup() {
	r.Engine = gin.Default()
	r.DB, _ = db.NewDB()
	h := handler.NewHTTPHandler(r.DB)
	userWeb := r.Engine.Group("/user")
	{
		userWeb.GET("/:id", h.GetUserProfile)
		userWeb.POST("/create", h.CreateNewUser)
	}
}
func NewRouter() Router {
	var r Router
	r.Setup()
	return r
}
