package routers

import (
	"github.com/bhbosman/golang/trading/ig/webserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
