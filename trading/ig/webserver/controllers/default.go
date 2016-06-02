package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "bhbosman"
	c.Data["Email"] = "bhbosman@gmail.com"
	c.TplName = "Main.tpl"
}
