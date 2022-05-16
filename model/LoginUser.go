package model

type LoginUser struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
