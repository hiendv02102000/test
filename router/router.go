package router

import (
	"test/db"
	"test/handler"
	"test/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	DB     db.Database
}

func (r *Router) Setup() error {
	err := r.DB.MigrationDB()
	if err != nil {
		return err
	}
	h := handler.NewHTTPHandler(r.DB)
	webAPI := r.Engine.Group("/web")
	{

		webAPI.POST("/login", h.Login)
		webAPI.POST("/register", h.RegisterUser)
	}
	userAPI := r.Engine.Group("/user")
	{
		userAPI.Use(middleware.AuthMiddleware(r.DB))
		{
			userAPI.GET("/:id", h.GetUserProfile)
			adminAPI := userAPI.Group("/admin")
			{
				adminAPI.Use(middleware.AuthAdminMiddleware())
				{
					adminAPI.DELETE("/delete", h.DeleteUser)
				}

			}
		}

	}
	return nil
}
func NewRouter() *Router {
	engine := gin.Default()
	db, err := db.NewDB()
	if err != nil {
		return nil
	}

	return &Router{Engine: engine, DB: db}
}
