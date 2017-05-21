package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
	"fmt"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	// 直接在屏幕上输出
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))

	this.TplName = "user.html"

	// 查询数据库
	var err error
	this.Data["Users"], err = models.GerAllUsers()
	if err != nil {
		beego.Error(err)
	}

	// 执行操作
	switch this.Input().Get("op") {
	case "add":
		uname := this.Input().Get("uname")
		pwd := this.Input().Get("pwd")
		// 存在
		if models.IsUserExist(uname) == nil {
			fmt.Println("已经存在")
			user := models.GetUserId(uname)
			err := models.UpdateUser(user.Id,uname,pwd)
			if err!=nil {
				beego.Error(err)
			}

			this.Ctx.Redirect(301,"/user")
			return
		}
		err := models.AddUser(uname,pwd)
		if err!=nil {
			beego.Error(err)
		}
		this.Ctx.Redirect(301,"/user")
	case "del":
		id := this.Input().Get("id")
		err := models.DelUser(id)
		if err!=nil {
			beego.Error(err)
		}
		this.Ctx.Redirect(301,"/user")
	case "login":
		uname := this.Input().Get("uname")
		pwd := this.Input().Get("pwd")
		autoLogin := this.Input().Get("autoLogin") == "on"
		if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
			maxAge := 0
			if autoLogin {
				maxAge = 1 << 31 -1
			}
			this.Ctx.SetCookie("uname",uname,maxAge,"/")
			this.Ctx.SetCookie("pwd",pwd,maxAge,"/")
			this.Ctx.Redirect(301,"/home")
		}
	}
	return
}

func (this *UserController) Post() {
	this.Get()
}