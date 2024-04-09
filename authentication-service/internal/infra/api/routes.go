package api

import (
	"github.com/deBeloper-code/authentication/internal/infra/api/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	// User
	user.RegisterRoutes(e)

}
