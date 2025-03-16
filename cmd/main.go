package main

import (
	"learn_ginmall/conf"
	"learn_ginmall/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	 r.Run(conf.HttpPort)
}