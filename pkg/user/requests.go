package user

type NewUserRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,gte=8"`
}
