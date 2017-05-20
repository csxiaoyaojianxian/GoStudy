package main

import (
	"beeblog/controllers"
	"beeblog/models"
	// "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	// 数据库调试
	orm.Debug = true
	// 自动建表，是否每次都自动建表，是否打印信息
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
