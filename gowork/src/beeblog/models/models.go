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
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Catagory struct {
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
	orm.RegisterModel(new(Catagory), new(Topic))
	// 注册驱动
	// orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	// 注册默认数据库，必须有一个数据库叫 default
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}
