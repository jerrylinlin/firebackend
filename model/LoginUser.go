package model

type LoginUser struct {
	ID       int64  `json:"id" gorm:"primary_key" form:"id"`
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (LoginUser) TableName() string {
	return "user"
}
