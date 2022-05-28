package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(r *gin.Engine, us UserService) {

	handlers := newUserHandlers(us)

	v1 := r.Group("/v1/users")
	{
		v1.GET("/", handlers.getUser)
		v1.POST("/signup", handlers.createUser)
	}
}
