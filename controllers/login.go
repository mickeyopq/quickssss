package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (自己 *LoginController) Get() {
	// 如果user按了登出....
	isExit := 自己.Input().Get("exit") // == "true"
	if isExit == "true" {
		自己.Ctx.SetCookie("uname", "", -1, "/")
		自己.Ctx.SetCookie("pwd", "", -1, "/")
		自己.Redirect("/", 301)
		return
	}
	自己.TplNames = "login.html"
	// c.Data["Condition11"] = false
}
func (自己 *LoginController) Post() {
	//登錄時驗証
	uname := 自己.Input().Get("uname")
	pwd := 自己.Input().Get("pwd")
	autoLogin := 自己.Input().Get("autoLogin") == "on"
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		自己.Ctx.SetCookie("uname", uname, maxAge, "/")
		自己.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		自己.Redirect("/", 301) //成功登錄就去那兒吧
		return

	} else {
		自己.Data["LgErr"] = true
	}
	自己.TplNames = "login.html"
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
