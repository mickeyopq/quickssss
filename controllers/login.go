package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (self *LoginController) Get() {
	// 判斷是否為登錄狀態,;如果user按了登出....
	isExit := self.Input().Get("exit") == "true"
	if isExit {
		self.Ctx.SetCookie("uname", "", -1, "/")
		self.Ctx.SetCookie("pwd", "", -1, "/")
		self.Redirect("/", 301)
		return
	}
	self.TplNames = "login.html"
	// c.Data["Condition11"] = false
}
func (self *LoginController) Post() {
	uname := self.Input().Get("uname")
	pwd := self.Input().Get("pwd")
	autoLogin := self.Input().Get("autoLogin") == "on"
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		self.Ctx.SetCookie("uname", uname, maxAge, "/")
		self.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		self.Redirect("/", 301) //成功登錄就去那兒吧
		return

	}
	// self.Redirect("/", 301) //成功登錄就去那兒吧
	// return
	self.TplNames = "login.html"

}

func 檢查帳號(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd

}
