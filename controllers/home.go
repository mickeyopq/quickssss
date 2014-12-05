package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (self *MainController) Get() {
	self.TplNames = "home.html"
	self.Data["IsHome"] = true
	self.Data["IsLogin"] = 檢查帳號(self.Ctx)
}
