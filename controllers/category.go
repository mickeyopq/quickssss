package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
	"quickssss/models"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

func (自己 *CategoryController) Get() {
	op := 自己.Input().Get("op")
	switch op {
	case "add":
		name := 自己.Input().Get("name")
		// ===DEBUG==============================
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error("===DEBUG==============================")
			beego.Error(err) //錯了就寫日記
		}
		自己.Redirect("/category", 302)
		return
	case "del":
		getid := 自己.Input().Get("id")
		id, _ := strconv.Atoi(getid)
		cid := int64(id)
		err := models.DelCategory(cid)
		if err != nil {
			beego.Error("===DEBUG==============================")
			beego.Error(err) //錯了就寫日記
		}
		// 自己.Redirect("/category", 302)
		// default:
		// 	return
	}

	// 	o := orm.NewOrm()
	// if num, err := o.Delete(&User{Id: 1}); err == nil {
	//     fmt.Println(num)
	// }
	自己.TplNames = "category.html"
	自己.Data["IsCategory"] = true
	var err error
	自己.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
