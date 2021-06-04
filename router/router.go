package router

import (
	"fmt"
	"test/db"
	"test/handler"
	"test/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	DB     db.Database
}

func (r *Router) Setup() {
	r.Engine = gin.Default()
	r.DB, _ = db.NewDB()
	//r.DB.MigrateDBWithGorm()
	err := r.DB.MigrateDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	h := handler.NewHTTPHandler(r.DB)
	appAPI := r.Engine.Group("/app")
	{
		userAPI := appAPI.Group("/user")
		{
			userAPI.Use(middleware.AuthClientMiddleware(r.DB))
			{
				userAPI.GET("/profile", h.GetUserProfile)

			}
			userAPI.Use(middleware.AuthAdminMiddleware(r.DB))
			{
				userAPI.DELETE("/delete", h.DeleteUser)
				userAPI.GET("/user_list", h.GetUserListProfile)
				userAPI.PATCH("/update", h.UpdateUser)
			}

		}

		webAPI := appAPI.Group("/web")
		{
			webAPI.POST("/login", h.Login)
			webAPI.POST("/create", h.CreateNewUser)
		}
	}
}
func NewRouter() Router {
	var r Router
	r.Setup()
	return r
}
