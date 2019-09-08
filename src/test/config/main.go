package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "./conf/config.conf")
	if err != nil {
		panic(err)
	}

	ip := conf.String("server::ip")
	fmt.Println("ip: ", ip)

}
