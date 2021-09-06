package middleware

import (
	"net/http"
	"strings"
	"test/db"
	"test/dto"
	"test/entity"
	"test/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetUserFromContext(c *gin.Context) entity.Users {
	value, exist := c.Get("user")
	if !exist {
		return entity.Users{}
	}
	return value.(entity.Users)
}
func AuthMiddleware(db db.Database) gin.HandlerFunc {

	return func(c *gin.Context) {

		clientToken := c.GetHeader("Authorization")
		if clientToken == "" {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  "Authorization Token is required",
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}
		extractedToken := strings.Split(clientToken, "Bearer ")
		clientToken = strings.TrimSpace(extractedToken[1])
		repo := repository.NewUserRepository(db)
		user, err := repo.FirstUser(entity.Users{
			Token: clientToken,
		})

		if gorm.IsRecordNotFoundError(err) {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  "Invalid Token",
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}
		if err != nil {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  err.Error(),
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}

		timeNow := time.Now()
		if timeNow.After(*user.TokenExpriedAt) {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  "Token Expired",
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
func AuthAdminMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		user := GetUserFromContext(c)
		if !user.Admin {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error:  "User isn't admin",
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}
		c.Next()
	}
}
