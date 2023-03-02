package main

import (
	"fmt"
	"seckill/config"
	"seckill/data"
	"seckill/router"
)

func main() {
	fmt.Println("---Starts server---")
	ds, err := data.InitData()
	if err != nil {
		fmt.Println("initData failed error:", err)
		return
	}
	r, err := router.InitRouter(ds)
	if err != nil {
		fmt.Println("initRouter failed error:", err)
	}
	conf, err := config.GetAppConfig()
	if err != nil {
		fmt.Println("getAppConfig failed error:", err)
	}
	r.Run(conf.HttpPort)

}
