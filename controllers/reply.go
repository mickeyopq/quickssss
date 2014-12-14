package controllers

import (
	"github.com/astaxie/beego"
	"quickssss/models"
)

type ReplyController struct {
	beego.Controller
}

func (自己 *ReplyController) Add() {
	tid := 自己.Input().Get("tid")
	err := models.AddReply(tid,
		自己.Input().Get("nickname"), 自己.Input().Get("content"))
	if err != nil {
		beego.Error(err)
	}

	自己.Redirect("/topic/view/"+tid, 301)
}

func (自己 *ReplyController) Del() {
	if !檢查帳號(自己.Ctx) {
		return
	}
	tid := 自己.Input().Get("tid")
	err := models.DelReply(自己.Input().Get("rid"))
	if err != nil {
		beego.Error(err)
	}

	自己.Redirect("/topic/view/"+tid, 301)
}
