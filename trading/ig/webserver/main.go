package main

import (
	_ "github.com/bhbosman/golang/trading/ig/webserver/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

