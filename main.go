package main

import (
	"./nickdblib"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

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

func HandleFindUserById(w http.ResponseWriter, req *http.Request) {
	strId := req.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Println(err)
	}
	user, err := nickdblib.FindUserById(id)

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
func HandleFindUserByName(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	user, err := nickdblib.FindUserByName(name)
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
func HandleUpdateUser(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("ReadAll error :", err)
		return
	}
	bodyStr := string(body)
	log.Println(bodyStr)
	var updateParam updateParam
	if err = json.Unmarshal(body, &updateParam); err == nil {

		err = nickdblib.UpdateUser(updateParam.Id, updateParam.Name, updateParam.Password, updateParam.Tag)
		if err != nil {
			w.Write(genResponse(500, "HandleUpdateUser exception", ""))
			return
		}
		w.Write(genResponse(0, "HandleUpdateUser success", ""))
		return
	} else {
		log.Println(err)
		w.Write(genResponse(500, "HandleUpdateUser error", ""))
	}
}
func HandleAddUser(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("ReadAll error :", err)
		return
	}
	bodyStr := string(body)
	log.Println(bodyStr)
	var user nickdblib.User

	if err := json.Unmarshal(body, &user); err == nil {
		err = nickdblib.AddUser(user)
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
func HandleDelUser(w http.ResponseWriter, req *http.Request) {
	strId := req.URL.Query().Get("id")
	soft := req.URL.Query().Get("soft")
	log.Println("HandleDelUser :", strId, soft)
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Println(err)
	}
	err = nickdblib.DelUser(id, soft)
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

func main() {
	//0.建立数据库连接
	nickdblib.InitDB()
	defer nickdblib.DB.Close()
	//1.http请求
	http.HandleFunc("/users/findUserById", HandleFindUserById)
	http.HandleFunc("/users/findByName", HandleFindUserByName)
	http.HandleFunc("/users/updateUser", HandleUpdateUser)
	http.HandleFunc("/users/addUser", HandleAddUser)
	http.HandleFunc("/users/delUser", HandleDelUser)
	http.ListenAndServe(":8001", nil)

}
