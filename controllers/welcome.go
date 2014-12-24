package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
)

type WelcomeController struct {
	beego.Controller
}

func (自己 *WelcomeController) Get() {
	自己.TplNames = "welcome.html"
}
