/*
* @Author: sunshine
* @Date:   2017-05-11 22:03:29
* @Last Modified by:   sunshine
* @Last Modified time: 2017-05-12 01:27:47
 */

package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/csxiaoyaojianxian/com"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"time"
	"strconv"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type User struct {
	Id              int64 // 名称为Id，类型为intXX，ORM默认为主键
	Name            string `orm:"index"`
	Pwd         	string
}

type Category struct {
	Id              int64 // 名称为Id，类型为intXX，ORM默认为主键
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

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
	// 创建数据库文件
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	// 注册模型
	orm.RegisterModel(new(User), new(Category), new(Topic))
	// 注册驱动
	// orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	// 注册默认数据库，必须有一个数据库叫 default
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func IsUserExist(name string) error {
	o := orm.NewOrm()
	// 新建一个对象用于判断是否已经存在
	user := &User{Name:name}
	// 查询数据库
	qs := o.QueryTable("user")
	err := qs.Filter("Name",name).One(user)
	// 判断是否已经存在
	// 存在
	if err == nil{
		return nil
	}
	// 不存在
	return err
}

func AddUser(name string, pwd string) error {
	o := orm.NewOrm()
	// 新建一个对象
	user := &User{Name:name,Pwd:pwd}
	// 存在
	if IsUserExist(name) == nil {
		return nil
	}
	// 不存在
	_, err := o.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func DelUser(id string) error {
	uid, err := strconv.ParseInt(id,10,64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	user := &User{Id:uid}
	_,err = o.Delete(user)
	return err
}

func UpdateUser(id int64, name string, pwd string) error {
	o := orm.NewOrm()
	// 新建一个对象
	user := &User{Id:id,Name:name,Pwd:pwd}
	// 存在
	if IsUserExist(name) == nil {
		_, err := o.Update(user)
		if err != nil {
			return err
		}
		return nil
	}
	// 不存在
	return nil
}

func GetUserId(name string) *User{
	o := orm.NewOrm()
	// 新建一个对象用于判断是否已经存在
	user := &User{Name:name}
	// 查询数据库
	qs := o.QueryTable("user")
	err := qs.Filter("Name",name).One(user)
	// 判断是否已经存在
	// 存在
	if err == nil{
		return user
	}
	// 不存在
	return user
}

func GerAllUsers() ([]*User,error) {
	o := orm.NewOrm()
	users := make([]*User,0)
	qs := o.QueryTable("user")
	_,err := qs.All(&users)
	return users,err
}