package user

import (
	"github.com/deBeloper-code/authentication/internal/infra/repositories"
	"github.com/deBeloper-code/authentication/internal/pkg/service/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	repo := repositories.NewClient()
	service := user.NewService(repo)
	handler := newHandler(service)
	v1 := e.Group("/api/v1")
	// Get users
	v1.GET("/users", handler.GetAll)
	v1.POST("user/authenticate", handler.LoginUser)
	v1.GET("/user/email", handler.GetUserByEmail)
	v1.GET("/user/id", handler.GetUserById)
	v1.POST("/user/id", handler.UpdateUser)
	v1.POST("/user", handler.CreateUser)
	v1.PUT("/user/id", handler.ResetPasswordUser)
	v1.DELETE("/user/id", handler.DeleteUser)
}
