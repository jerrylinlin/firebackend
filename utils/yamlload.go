package utils

import (
	"fmt"
	"github.com/jinzhu/configor"
)

var Cfg *DBConfig

func ConfigRead(configFile string) {
	config := new(DBConfig)
	config.readConfigFromYaml(configFile)
	fmt.Println("Final config: ", config)
	Cfg = config
	fmt.Println(Cfg.Password)
}

func (this *DBConfig) readConfigFromYaml(configFile string) {
	if err := configor.Load(this, configFile); err != nil {
		fmt.Println("read config yaml err: " + err.Error())
		return
	}
	fmt.Println("Config from yaml file: ", this)
}
