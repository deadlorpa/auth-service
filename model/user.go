package model

type User struct {
	Id       string `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	IdRole   string `json:"id_role"`
}

type UserSignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignInResponce struct {
	Id       string `json:"id_user"`
	Username string `json:"username"`
	Role     Role   `json:"role"`
	Token    string `json:"token"`
}
