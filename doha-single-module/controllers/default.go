package controllers

import (
	"doha-single-module/business/service"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
	UserService *service.UserService
}

func (c *MainController) Get() {
	users, err := c.UserService.GetAll()
	if err != nil {
		c.Data["users"] = nil
	} else {
		c.Data["users"] = users
	}
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
