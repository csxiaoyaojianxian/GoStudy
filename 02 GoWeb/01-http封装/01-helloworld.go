/*
* @Author: csxiaoyao
* @Date:   2017-05-11 00:34:25
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-11 00:48:07
 */

package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/", &HomeController{})
	beego.Run()
}
