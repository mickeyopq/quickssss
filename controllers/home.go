package controllers

import (
	"github.com/astaxie/beego"
	"quickssss/models"
)

type MainController struct {
	beego.Controller
}

func (自己 *MainController) Get() {
	自己.TplNames = "home.html"
	自己.Data["IsHome"] = true
	自己.Data["IsLogin"] = 檢查帳號(自己.Ctx)

	//叫model裡的func傳文章
	topics, err := models.GetAllTopics(自己.Input().Get("cate"), true) //是否倒排序
	if err != nil {
		beego.Error(err)
	}
	自己.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	自己.Data["Categories"] = categories

}
