package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
)

var insertId int

func TestAdd(t *testing.T) {
	var user User
	user.Name = "test"
	user.Password = "testPassword"
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.Post("http://localhost:8001/users/addUser", "application/json;charset=utf-8", bytes.NewBuffer([]byte(jsonBytes)))
	if err != nil {
		t.Fatal("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Fatal error ", err.Error())
	}

	var rd returnData
	if err := json.Unmarshal(content, &rd); err == nil {
		if rd.Code != 0 {
			log.Println("insert test error :", rd)
			t.Fatal("addUser error ", rd)
		}
	} else {
		t.Fatal("json Fatal error ", err.Error())
	}
}
func TestFind(t *testing.T) {
	name := "test"
	//1.根据姓名查找
	res, err := http.Get("http://localhost:8001/users/findByName?name=" + name)
	if err != nil {
		t.Fatal("findByName Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("findByName Fatal error ", err.Error())
	}

	var user User
	var rd returnData
	if err := json.Unmarshal(content, &rd); err == nil {
		if rd.Code != 0 {
			log.Println("findByName test error :", rd)
			t.Fatal("findByName error ", rd)
		}
		if err := json.Unmarshal([]byte(rd.Data), &user); err == nil {
		} else {
			t.Fatal("findByName UnMarshal error ", err)
		}
	} else {
		t.Fatal("findByName json Fatal error ", err.Error())
	}

	log.Println("findByName user is :", user)
	insertId = user.Id
	//2.根据id查找  (上面根据name查出一个用户）
	res, err = http.Get("http://localhost:8001/users/findById?id=" + strconv.Itoa(insertId))
	if err != nil {
		t.Fatal("findById Fatal error ", err.Error())
	}

	content, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("findById Fatal error ", err.Error())
	}
	if err := json.Unmarshal(content, &rd); err == nil {
		if rd.Code != 0 {
			log.Println("findById test error :", rd)
			t.Fatal("findById error ", rd)
		}
		if err := json.Unmarshal(content, &user); err == nil {
		} else {
			t.Fatal("findById UnMarshal error ", err)
		}
	} else {
		t.Fatal("findById json Fatal error ", err.Error())
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
	jsonBytes, err := json.Marshal(updateParam)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.Post("http://localhost:8001/users/updateUser", "application/json;charset=utf-8", bytes.NewBuffer([]byte(jsonBytes)))
	if err != nil {
		t.Fatal("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Update Fatal error ", err.Error())
	}

	var rd returnData
	if err := json.Unmarshal(content, &rd); err == nil {
		if rd.Code != 0 {
			log.Println("Update test error :", rd)
			t.Fatal("Update error ", rd)
		}
	} else {
		t.Fatal("Update json Fatal error ", err.Error())
	}

	//2.再修改密码
	updateParam.Tag = 2
	jsonBytes, err = json.Marshal(updateParam)
	if err != nil {
		t.Fatal(err)
	}
	res, err = http.Post("http://localhost:8001/users/updateUser", "application/json;charset=utf-8", bytes.NewBuffer([]byte(jsonBytes)))
	if err != nil {
		t.Fatal("Fatal error ", err.Error())
	}

	content, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Update Fatal error ", err.Error())
	}
	if err := json.Unmarshal(content, &rd); err == nil {
		if rd.Code != 0 {
			log.Println("Update test error :", rd)
			t.Fatal("Update error ", rd)
		}
	} else {
		t.Fatal("Update json Fatal error ", err.Error())
	}

	var user User
	//3.查出id=38这个用户  看看账号密码变没变
	res, err = http.Get("http://localhost:8001/users/findById?id=" + strconv.Itoa(updateParam.Id))
	if err != nil {
		t.Fatal("findById Fatal error ", err.Error())
	}

	content, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("findById Fatal error ", err.Error())
	}
	if err := json.Unmarshal(content, &rd); err == nil {
		if rd.Code != 0 {
			log.Println("findById test error :", rd)
			t.Fatal("findById error ", rd)
		}
		if err := json.Unmarshal([]byte(rd.Data), &user); err == nil {
		} else {
			t.Fatal("findById UnMarshal error ", err)
		}
	} else {
		t.Fatal("findById json Fatal error ", err.Error())
	}

	log.Println("user is :", user)
	log.Println("updateParam is :", updateParam)
	//看下目前的姓名和密码是不是刚刚更新的
	if user.Name != updateParam.Name || user.Password != updateParam.Password {
		t.Fatal("update ERROR")
	}
}

func TestDelete(t *testing.T) {
	//1.根据姓名查找
	res, err := http.Get("http://localhost:8001/users/delUser?id=" + strconv.Itoa(insertId))
	if err != nil {
		t.Fatal("findByName Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("delete Fatal error ", err.Error())
	}

	var rd returnData
	if err := json.Unmarshal(content, &rd); err == nil {
		if rd.Code != 0 {
			log.Println("delete test error :", rd)
			t.Fatal("delete error ", rd)
		}
	} else {
		t.Fatal("delete json Fatal error ", err.Error())
	}
}
