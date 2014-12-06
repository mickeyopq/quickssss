package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"quickssss/models"
	_ "quickssss/routers"
)

func init() {
	// 注冊數據庫
	models.RegisterDB()
}
func main() {
	// 開啟 ORM 調試模式
	orm.Debug = true
	// 自動建表
	orm.RunSyncdb("default", false, true)
	// 注册 beego 路由
	// beego.Router("/", &controllers.HomeController{})
	// orm.DebugLog = orm.NewLog(w)
	beego.Error("\n===DEBUG==============================")
	beego.Run()
}
