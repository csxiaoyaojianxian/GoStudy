package controllers

import (
	"github.com/astaxie/beego"
	// "strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"

	c.Ctx.WriteString("appname:" + beego.AppConfig.String("appname") +
		"\nhttpport:" + beego.AppConfig.String("httpport") +
		"\nrunmode:" + beego.AppConfig.String("runmode"))

	// hp := strconv.Itoa(beego.HttpPort)
	// c.Ctx.WriteString("\n\nappname:" + beego.AppName +
	// 	"\nhttpport:" + hp +
	// 	"\nrunmode:" + beego.RunMode)

	beego.Trace("trace test1")
	beego.Info("info test1")
}
