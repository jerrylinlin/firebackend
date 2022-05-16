package model

import "github.com/jinzhu/gorm"

type UserReg struct {
	gorm.Model
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Email    string `json:"email" yaml:"email"`
	Phone    string `json:"phone" yaml:"phone"`
	Status   int    `json:"status" yaml:"status"`
	Created  int64  `json:"created" yaml:"created"`
	Updated  int64  `json:"updated" yaml:"updated"`
}
