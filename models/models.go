package models

import (
	"os"
	"path"
	"time"

	"github.com/Unknwon/com"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	// "設置數據庫路徑"
	_DB_NAME = "data/beeblog.db"
	// 設置數據庫名稱
	_SQLITE3_DRIVER = "sqlite3"
)

// 分類
type Category struct {
	Id    int64 `orm:"auto"`
	Title string
	// Created         time.Time `orm:"index"`
	// Views           int64
	// TopicTime       time.Time `orm:"index"`
	// TopicCount      int64
	// TopicLastUserId int64
}

// 文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	// 檢查數據庫文件
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

// AddMod_ppp insert a new Mod_ppp into database and returns
// last inserted Id on success.
// func AddMod_ppp(m *Mod_ppp) (id int64, err error) {
// 	o := orm.NewOrm()
// 	id, err = o.Insert(m)
// 	return
// }
func DelCategory(id int64) error {
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&Category{Id: id})
	return nil
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	o.Using("default")
	tb := new(Category)
	tb.Title = name
	o.Insert(tb)
	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}
