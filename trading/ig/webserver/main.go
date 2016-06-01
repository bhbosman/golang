package main

import (
	"github.com/astaxie/beego"
	_ "github.com/bhbosman/golang/trading/ig/webserver/routers"
)

func main() {
	beego.Run()

}
