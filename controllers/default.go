package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "mememe"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Name"] = "hunter.yang"
	c.TplName = "mine.tpl"
}
