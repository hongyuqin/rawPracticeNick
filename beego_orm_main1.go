package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)
import _ "github.com/go-sql-driver/mysql"

type User1 struct {
	Id   int
	Name string
}

//自定义表名
func (u *User1) TableName() string {
	return "auth_user"
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:hongyuqin@/orm_test?charset=utf8")

	orm.RegisterModel(new(User1))

	orm.RunSyncdb("default", false, true)
}

func main() {
	//o := orm.NewOrm()
	//o.Using("default")
	logs.Debug("this is debug :", 2019)
}
