package common

import (
	"firebackend/model"
	"firebackend/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"log"
)

func MysqlInit() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.Cfg.User, utils.Cfg.Password, utils.Cfg.Host, utils.Cfg.Port, utils.Cfg.DBName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&model.UserReg{})
	return db
}
