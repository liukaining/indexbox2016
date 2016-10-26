package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "indexbox.me"
	c.Data["Email"] = "liukaining2016@163.com"
	c.TplName = "index.tpl"
}
