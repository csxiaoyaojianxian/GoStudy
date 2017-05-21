package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type HomeController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// 变量定义 map
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	// 布尔值
	c.Data["TrueCond"] = true
	c.Data["FalseCond"] = false

	// 定义结构体
	type u struct {
		Name string
		Age  int
		Sex  string
	}
	user := &u{
		Name: "csxiaoyao",
		Age:  20,
		Sex:  "Male",
	}
	// 字段 User 大小写效果相同
	c.Data["User"] = user

	// 数组
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	c.Data["Nums"] = nums

	// 模板变量测试数据
	c.Data["TplVar"] = "sunshine"

	// 模板函数
	c.Data["Html"] = "<div>hello beego</div>"

	// 字符转义
	c.Data["Pipe"] = "<div>hello beego</div>"

	// 模板嵌套

}

func (c *HomeController) Get() {
	c.TplName = "home.html"
	// 变量定义 map
	c.Data["Website"] = "www.csxiaoyao.com"
	c.Data["Email"] = "sunjianfeng@csxiaoyao.com"
}