package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//用户信息
type User struct {
	Id       int
	Name     string
	Password string
	Status   int
}

//接口返回参数
type returnData struct {
	Code int
	Msg  string
	Data string
}

//更新用户参数
type updateParam struct {
	Id       int
	Name     string
	Password string
	Tag      int
}

//1.1 查找：根据id查找一个用户
func findById(id int) *User {
	log.Println("findById : ", id)
	row := db.QueryRow("select id,password,status,name from users where id = ?", id)
	var user User
	var name, password string
	var status int
	err := row.Scan(&id, &password, &status, &name)
	if err != nil {
		log.Println("search error ", err)
		return nil
	}
	user.Id = id
	user.Password = password
	user.Status = status
	user.Name = name
	log.Println(user)
	return &user
}

//1.2 查找：根据姓名查找用户
func findByName(searchName string) *User {
	log.Println("findByName : ", searchName)
	row := db.QueryRow("select id,password,status,name from users where name = ?", searchName)
	var user User
	var name, password string
	var status, id int
	err := row.Scan(&id, &password, &status, &name)
	if err != nil {
		log.Println("search error ", err)
		return nil
	}
	user.Id = id
	user.Password = password
	user.Status = status
	user.Name = name
	log.Println(user)
	return &user
}

//2.新增用户
func addUser(user User) error {
	log.Println("addUser : ", user)
	result, err := db.Exec(
		"INSERT INTO users (name, password,status) VALUES (?, ?,?)",
		user.Name,
		user.Password,
		1,
	)
	printResult(result, err)
	return err
}

//3.删除用户
func delUser(id int, soft string) error {
	log.Println("delUser : ", id, soft)
	//1.软删除
	if soft == "1" {
		result, err := db.Exec("UPdate users set status=0 where id=? ", id)
		affectNum, err := result.RowsAffected()
		log.Println("affectNum err ", affectNum, err)
		if affectNum == 0 {
			return errors.New("no update")
		}
		printResult(result, err)
		return err
	}
	//2.硬删除
	result, err := db.Exec("DELETE FROM users WHERE id=?", id)
	printResult(result, err)
	return err
}

//4.1 修改用户姓名、密码(根据id)
//tag=1修改姓名  tag=2修改密码
func updateUser(id int, name string, password string, tag int) error {
	log.Println("update User : ", id, name, password, tag)
	if tag == 1 {
		result, err := db.Exec("UPdate users set name=? where id=? ", name, id)
		printResult(result, err)
		if err != nil {
			return err
		}
		return nil
	}
	result, err := db.Exec("UPdate users set password=? where id=? ", password, id)
	printResult(result, err)
	if err != nil {
		return err
	}
	affectNum, err := result.RowsAffected()
	log.Println("affectNum err ", affectNum, err)
	if affectNum == 0 {
		return errors.New("no affect")
	}
	return nil
}

func printResult(result sql.Result, err error) {
	if err != nil {
		log.Println("错误信息为：", err)
	}
}

func findByIdHttp(w http.ResponseWriter, req *http.Request) {
	strId := req.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Println(err)
	}
	user := findById(id)

	if user == nil {
		w.Write(genResponse(500, "customer not found", ""))
		return
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jsonBytes))
	w.Write(genResponse(0, "", string(jsonBytes)))
}
func findByNameHttp(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	user := findByName(name)
	if user == nil {
		w.Write(genResponse(500, "customer not found", ""))
		return
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.Write(genResponse(500, err.Error(), ""))
		return
	}
	log.Println(string(jsonBytes))
	w.Write(genResponse(0, "", string(jsonBytes)))
}
func updateUserHttp(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	body_str := string(body)
	log.Println(body_str)
	var updateParam updateParam
	if err := json.Unmarshal(body, &updateParam); err == nil {

		err = updateUser(updateParam.Id, updateParam.Name, updateParam.Password, updateParam.Tag)
		if err != nil {
			w.Write(genResponse(500, "updateUserHttp exception", ""))
			return
		}
		w.Write(genResponse(0, "updateUserHttp success", ""))
		return
	} else {
		log.Println(err)
		w.Write(genResponse(500, "updateUserHttp error", ""))
	}
}
func addUserHttp(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	body_str := string(body)
	log.Println(body_str)
	var user User

	if err := json.Unmarshal(body, &user); err == nil {
		err = addUser(user)
		if err != nil {
			w.Write(genResponse(500, "insert exception", ""))
			return
		}
		w.Write(genResponse(0, "insert success", ""))
		return
	} else {
		log.Println(err)
		w.Write(genResponse(500, "insert error", ""))
	}
}
func delUserHttp(w http.ResponseWriter, req *http.Request) {
	strId := req.URL.Query().Get("id")
	soft := req.URL.Query().Get("soft")
	log.Println("delUserHttp :", strId, soft)
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Println(err)
	}
	err = delUser(id, soft)
	if err != nil {
		w.Write(genResponse(500, err.Error(), ""))
		return
	}
	w.Write(genResponse(0, "del success", ""))
	return
}

func genResponse(code int, msg string, data string) []byte {
	log.Println("genResponse :", code, msg, data)
	returnData := returnData{code, msg, data}
	jsonBytes, err := json.Marshal(returnData)
	if err != nil {
		log.Println("genResponse  error :", err)
	}
	return jsonBytes
}

//全局变量 数据库连接
var db *sql.DB

func main() {
	//0.建立数据库连接
	dbMain, err := sql.Open("mysql", "root:hongyuqin@/mydb")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer dbMain.Close()
	err = dbMain.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successful connect to mysql")
	db = dbMain
	//1.http请求
	http.HandleFunc("/users/findById", findByIdHttp)
	http.HandleFunc("/users/findByName", findByNameHttp)
	http.HandleFunc("/users/updateUser", updateUserHttp)
	http.HandleFunc("/users/addUser", addUserHttp)
	http.HandleFunc("/users/delUser", delUserHttp)
	http.ListenAndServe(":8001", nil)

}
