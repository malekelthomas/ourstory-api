package user

import (
	"github.com/gin-gonic/gin"
	"github.com/malekelthomas/ourstory-api/server"
)

type handlers struct {
	us UserService
}

func newUserHandlers(us UserService) *handlers {
	return &handlers{us: us}
}

func (h *handlers) getUser(c *gin.Context) {
	id := c.DefaultQuery("id", "none")
	user, err := h.us.Get(id)
	if err != nil {
		c.JSON(500, server.ServerErrorMsg{Message: "user not found", Error: err.Error()})
	} else {
		c.JSON(200, user)
	}
}

func (h *handlers) createUser(c *gin.Context) {
	var userJSON NewUserRequest
	if err := c.BindJSON(&userJSON); err != nil {
		c.JSON(400, gin.H{"Message": "malformed request,", "Error": err.Error()})
	} else {
		user, err := h.us.Create(userJSON.UserName, userJSON.Password, UserRole(0))
		if err != nil {
			c.JSON(500, server.ServerErrorMsg{Message: "unable to create user", Error: err.Error()})
		} else {
			c.JSON(200, user)
		}
	}
}
