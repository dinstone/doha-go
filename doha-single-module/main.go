package main

import (
	_ "doha-single-module/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

