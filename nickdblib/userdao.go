package nickdblib

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

//用户信息
type User struct {
	Id       int
	Name     string
	Password string
	Status   int
}

//全局变量 数据库连接
var DB *sql.DB

func InitDB() {
	dbMain, err := sql.Open("mysql", "root:hongyuqin@/klerpdb")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}

	err = dbMain.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successful connect to mysql")
	DB = dbMain
}

//1.1 查找：根据id查找一个用户
func FindUserById(id int) (*User, error) {
	log.Println("FindUserById : ", id)
	row := DB.QueryRow(`select id,IFNULL(password,""),IFNULL(status,0),IFNULL(name,"") from users where id = ?`, id)
	var user User
	var name, password string
	var status int
	err := row.Scan(&id, &password, &status, &name)
	if err != nil {
		log.Println("search error ", err)
		return nil, err
	}
	user.Id = id
	user.Password = password
	user.Status = status
	user.Name = name
	log.Println(user)
	return &user, nil
}

//1.2 查找：根据姓名查找用户
func FindUserByName(searchName string) (*User, error) {
	log.Println("FindUserByName : ", searchName)
	likeName := fmt.Sprintf("n%%%s%%", searchName)
	row := DB.QueryRow(`select id,password,status,name from users where name like ?`, likeName)
	var user User
	var name, password string
	var status, id int
	err := row.Scan(&id, &password, &status, &name)
	if err != nil {
		log.Println("search error ", err)
		return nil, err
	}
	user.Id = id
	user.Password = password
	user.Status = status
	user.Name = name
	log.Println(user)
	return &user, nil
}

//2.新增用户
func AddUser(user User) error {
	log.Println("addUser : ", user)
	result, err := DB.Exec(
		"INSERT INTO users (name, password,status) VALUES (?, ?,?)",
		user.Name,
		user.Password,
		1,
	)
	printResult(result, err)
	return err
}

//3.删除用户
func DelUser(id int, soft string) error {
	log.Println("delUser : ", id, soft)
	//1.软删除
	if soft == "1" {
		result, err := DB.Exec("UPdate users set status=0 where id=? ", id)
		affectNum, err := result.RowsAffected()
		log.Println("affectNum err ", affectNum, err)
		if affectNum == 0 {
			return errors.New("no update")
		}
		printResult(result, err)
		return err
	}
	//2.硬删除
	result, err := DB.Exec("DELETE FROM users WHERE id=?", id)
	printResult(result, err)
	return err
}

//4.1 修改用户姓名、密码(根据id)
//tag=1修改姓名  tag=2修改密码
func UpdateUser(id int, name string, password string, tag int) error {
	log.Println("update User : ", id, name, password, tag)
	if tag == 1 {
		result, err := DB.Exec("UPdate users set name=? where id=? ", name, id)
		printResult(result, err)
		if err != nil {
			return err
		}
		return nil
	}
	result, err := DB.Exec("UPdate users set password=? where id=? ", password, id)
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
