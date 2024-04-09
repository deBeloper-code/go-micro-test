package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deBeloper-code/authentication/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService ports.UserService
}

func newHandler(service ports.UserService) *userHandler {
	return &userHandler{
		userService: service,
	}
}

func (u *userHandler) GetAll(c *gin.Context) {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, users)
}

type EmailRequest struct {
	Email string `json:"email" binding:"required"`
}

func (u *userHandler) GetUserByEmail(c *gin.Context) {
	var emailReq EmailRequest
	// Validate Body
	if err := c.BindJSON(&emailReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid!"})
		return
	}
	// Send service
	user, err := u.userService.GetUserByEmail(emailReq.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user finds"})
		return
	}

	c.JSON(http.StatusOK, user)
}

type IdRequest struct {
	ID int `json:"id" binding:"required"`
}

func (u *userHandler) GetUserById(c *gin.Context) {
	var idReq IdRequest
	// Validate Body
	if err := c.BindJSON(&idReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid!"})
		return
	}
	// Send service
	user, err := u.userService.GetUserById(idReq.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user finds"})
		return
	}

	c.JSON(http.StatusOK, user)
}

type UserInfoRequest struct {
	ID        int       `json:"id" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Active    int       `json:"active" binding:"required"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *userHandler) UpdateUser(c *gin.Context) {
	var userInfo UserInfoRequest
	// Validate Body
	if err := c.BindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Send service
	user, err := u.userService.UpdateUserInfo(userInfo.ID, userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User was updated!", "userId": user.ID})
}
