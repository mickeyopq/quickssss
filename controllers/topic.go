package controllers

import (
	"github.com/astaxie/beego"
	"quickssss/models"
)

// oprations for Topic
type TopicController struct {
	beego.Controller
}

func (自己 *TopicController) Get() {
	自己.Data["IsTopic"] = true
	自己.TplNames = "topic.html"
	自己.Data["IsLogin"] = 檢查帳號(自己.Ctx)

	topics, err := models.GetAllTopics("", true) //是否倒排序
	if err != nil {
		beego.Error(err)
	}
	自己.Data["Topics"] = topics
}

func (自己 *TopicController) Post() {
	// 自己.Ctx.WriteString("幹")
	if !檢查帳號(自己.Ctx) {
		自己.Redirect("/login", 302)
		return
	}
	//接收文章表單
	// 解析表单
	tid := 自己.Input().Get("tid")
	title := 自己.Input().Get("title")
	category := 自己.Input().Get("category")
	content := 自己.Input().Get("content")
	// category := 自己.Input().Get("category")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, content, category)
		// err = models.AddTopic(title, category, content)
	} else {
		err = models.ModifyTopic(tid, title, content, category)
	}
	if err != nil {
		beego.Error(err)
	}
	自己.Redirect("/topic", 302)
}

func (自己 *TopicController) View() {
	自己.TplNames = "topic_view.html"
	//Params的0是網址的第一個..../web/0/1/2/3
	tid := 自己.Ctx.Input.Param("0")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		自己.Redirect("/", 302)
		return
	}
	自己.Data["Topic"] = topic
	replies, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
		return
	}
	自己.Data["Replies"] = replies
	// 自己.Data["IsLogin"] = 檢查帳號(自己.Ctx)
	// 自己.Data["IsLogin"] = 檢查帳號(自己.Ctx)
	// 自己.Data["tid"] = tid
	// 自己.Data["Category"] = "Category還沒定義"
}
func (自己 *TopicController) Del() {
	if !檢查帳號(自己.Ctx) {
		自己.Ctx.WriteString("不可以這樣")
		return
	}
	err := models.DelTopic(自己.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	自己.Redirect("/topic", 302)
}

func (自己 *TopicController) Modify() {
	if !檢查帳號(自己.Ctx) {
		自己.Ctx.WriteString("不可以這樣")
		return
	}
	自己.TplNames = "topic_modify.html"
	tid := 自己.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		自己.Redirect("/", 302)
		return
	}
	自己.Data["Topic"] = topic
	自己.Data["Tid"] = tid
	自己.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

func (自己 *TopicController) Add() {
	if 檢查帳號(自己.Ctx) {
		自己.TplNames = "topic_add.html"
		自己.Data["IsLogin"] = true
		var err error
		自己.Data["Categories"], err = models.GetAllCategories()
		if err != nil {
			beego.Error(err)
		}
	} else {
		自己.Ctx.WriteString("不可以這樣地新增.....")
		return
	}
}
