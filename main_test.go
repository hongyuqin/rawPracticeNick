package main

import (
	"./nickdblib"
	"log"
	"testing"
)

var insertId int

func TestAdd(t *testing.T) {

	var user nickdblib.User
	user.Name = "test"
	user.Password = "testPassword"
	err := nickdblib.AddUser(user)
	if err != nil {
		t.Fatal("nickdblib.AddUser Fatal error ", err.Error())
	}
}
func TestFind(t *testing.T) {
	name := "test"
	//1.根据姓名查找
	user, err := nickdblib.FindUserByName(name)
	if err != nil {
		t.Fatal("findByName Fatal error ", err.Error())
	}

	insertId = user.Id
	//2.根据id查找  (上面根据name查出一个用户）
	user, err = nickdblib.FindUserById(insertId)
	if err != nil {
		t.Fatal("findById Fatal error ", err.Error())
	}
	if err != nil {
		t.Fatal("findById Fatal error ", err.Error())
	}

}

func TestUpdate(t *testing.T) {

	//测试更新用户  根据id 把名字更改为  test1, 密码更改为12345
	//1.先修改用户名
	var updateParam updateParam
	updateParam.Id = insertId
	updateParam.Name = "test1"
	updateParam.Password = "12345"
	updateParam.Tag = 1
	err := nickdblib.UpdateUser(insertId, updateParam.Name, updateParam.Password, 1)

	if err != nil {
		t.Fatal("updateUser1  ", err)
	}
	err = nickdblib.UpdateUser(insertId, updateParam.Name, updateParam.Password, 2)
	if err != nil {
		t.Fatal("Fatal error ", err.Error())
	}

	user, err := nickdblib.FindUserById(insertId)
	//看下目前的姓名和密码是不是刚刚更新的
	if user.Name != updateParam.Name || user.Password != updateParam.Password {
		t.Fatal("update ERROR")
	}
}

func TestDelete(t *testing.T) {
	//1.根据姓名查找
	err := nickdblib.DelUser(insertId, "0")
	if err != nil {
		t.Fatal("delete Fatal error ", err.Error())
	}
}

func TestMain(m *testing.M) {
	log.Println("begin")
	nickdblib.InitDB()
	defer nickdblib.DB.Close()
	m.Run()
	log.Println("end")
}
