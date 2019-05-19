package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

/**
CREATE TABLE IF NOT EXISTS `mydb`.`users` (
	  `id` INT NOT NULL AUTO_INCREMENT,
	  `name` VARCHAR(45) NOT NULL COMMENT '用户名',
	  `password` VARCHAR(45) NOT NULL COMMENT '密码',
	  `status` INT NOT NULL DEFAULT 0 COMMENT '0:有效帐户 1:无效帐户',
	  PRIMARY KEY (`id`),
	  UNIQUE INDEX `idx_user_01` (`name` ASC))
	ENGINE = InnoDB
	DEFAULT CHARACTER SET = utf8
	COMMENT = '用户表'
*/
type User struct {
	id       int
	name     string
	password string
	status   int
}

func SayHello(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Hello"))
}

//根据id查找一个用户
func findById(db *sql.DB, id int) User {
	rows, err := db.Query("select id,password,status,name from users where id = ?", id)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var user User
	var name, password string
	var status int
	for rows.Next() {
		err := rows.Scan(&id, &password, &status, &name)
		if err != nil {
			log.Fatal(err)
		}
		user.id = id
		user.password = password
		user.status = status
		user.name = name
		log.Println(name)
	}
	return user
}

//根据姓名查找用户
func findByName(db *sql.DB, searchName string) User {
	rows, err := db.Query("select id,password,status,name from users where name = ?", searchName)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var user User
	var name, password string
	var status, id int
	for rows.Next() {
		err := rows.Scan(&id, &password, &status, &name)
		if err != nil {
			log.Fatal(err)
		}
		user.id = id
		user.password = password
		user.status = status
		user.name = name
	}
	return user
}
func main() {
	/*http.HandleFunc("/findOne", SayHello)
	http.ListenAndServe(":8001", nil)
	*/
	db, err := sql.Open("mysql", "root:hongyuqin@/mydb")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	user := findById(db, 1)
	fmt.Println(user)

	user = findByName(db, "guotie1")
	fmt.Println(user)
	/*users := findMulti(db,"guotie")
	fmt.Println(users)*/

}
