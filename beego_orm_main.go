package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:hongyuqin@/orm_test?charset=utf8")
}

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	// 需要在init中注册定义的model
	beego.Notice("hhaha init")
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))
	//建表
	//orm.RunSyncdb("default", false, true)
}

func read() {
	o := orm.NewOrm()
	user := User{Name: "slene"}
	err := o.Read(&user, "Name")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Name)
	}
}

//原生sql查询
//1.更新
func exec(o orm.Ormer) {
	res, err := o.Raw("update user set name = ?", "your").Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums :", num)
	}
}

//2.查找某一条数据
func findOne(o orm.Ormer) {
	var user User
	err := o.Raw("select id,name from user where id = ?", 1).QueryRow(&user)
	if err != nil {
		fmt.Printf("error is :%s", err)
		return
	}
	fmt.Println("findOne result :", user)
}

//3.查找多条数据
func findMulti(o orm.Ormer) {
	var users []User
	nums, err := o.Raw("select id,name from user where name = ?", "your").QueryRows(&users)
	if err != nil {
		fmt.Println("findMulti error :", err)
		return
	}
	fmt.Println("findMulti result :", nums, users)
}

//4.rowsToStruct
type Options struct {
	Total int
	Found int
}

func testRowsToStruct(o orm.Ormer) {
	res := new(Options)
	_, err := o.Raw("select id,name from user ").RowsToStruct(res, "id", "name")
	if err != nil {
		fmt.Println("testRowsToStruct error :", err)
	} else {
		fmt.Println(res.Total)
		fmt.Println(res.Found)
	}
}
func main() {
	/*orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))*/
	//read()
	o := orm.NewOrm()
	//exec(o)
	//findOne(o)
	//findMulti(o)
	testRowsToStruct(o)
}
