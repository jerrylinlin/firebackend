package main

import (
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
	R := route.RoutRegister()
	panic(R.Run(":8000"))

}
