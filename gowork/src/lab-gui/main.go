package main

import (
	_ "lab-gui/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lab-gui/controllers"
	"lab-gui/models"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{})
	beego.Run()
}

