package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/csxiaoyaojianxian/com"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
)

const (
	_DB_NAME        = "data/menu.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Menu struct {
	Id              int64
	Cmd             string `orm:"index"`
	Desc         	string
	Handler		string
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Menu))
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func IsCmdExist(cmd string) error {
	o := orm.NewOrm()
	menu := &Menu{Cmd:cmd}
	qs := o.QueryTable("menu")
	err := qs.Filter("Menu",menu).One(menu)
	// 存在
	if err == nil{
		return nil
	}
	// 不存在
	return err
}

func AddCmd(cmd string,desc string) error {
	o := orm.NewOrm()
	menu := &Menu{Cmd:cmd,Desc:desc}
	// 存在
	if IsCmdExist(cmd) == nil {
		return nil
	}
	// 不存在
	_, err := o.Insert(menu)
	if err != nil {
		return err
	}
	return nil
}

func DelCmd(id string) error {
	cid, err := strconv.ParseInt(id,10,64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	menu := &Menu{Id:cid}
	_,err = o.Delete(menu)
	return err
}

func GetCmdId(cmd string) *Menu{
	o := orm.NewOrm()
	menu := &Menu{Cmd:cmd}
	qs := o.QueryTable("menu")
	qs.Filter("Cmd",cmd).One(menu)
	return menu
}

func GerAllCmds() ([]*Menu,error) {
	o := orm.NewOrm()
	menus := make([]*Menu,0)
	qs := o.QueryTable("menu")
	_,err := qs.All(&menus)
	return menus,err
}