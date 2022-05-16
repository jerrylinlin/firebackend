package main

import (
	"firebackend/common"
	"firebackend/route"
	"firebackend/utils"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	var configFile = flag.String("configFile", "config/config.yml", "配置文件路径")
	flag.Parse()
	utils.ConfigRead(*configFile)
	db := common.MysqlInit()
	R := route.RoutRegister()
	panic(R.Run(":8080"))
	defer db.Close()

}
