package handlers

import(
	"net/http"
	"bytes"
	"encoding/json"
	"strconv"
	"log"
	"fmt"
	"main/module"
)

func GetTodos (w http.ResponseWriter, r *http.Request){
	//requestURLクエリパラメータから渡されたuser_idを取得する（これはstring型で取得される）
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == ""{
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	//上記で取得したuserIDStrをintに変換
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}
	
	//usersテーブルからuserを取得
	user, err := module.GetUser(userID)
	if err != nil{
		http.Error(w, "the user with this user_id is not exist", http.StatusBadRequest)
	}

	//上記で取得したuserに対応するtodoをtodosとして取得
	todos, err := user.GetTodosByUserID()
	if err != nil{
		http.Error(w, "todos with this user_id is not exist", http.StatusBadRequest)
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&todos); err != nil{
		log.Fatalln(err)
	}
	fmt.Println(buf.String())

	_, err = fmt.Fprint(w, buf.String())
	if err != nil{
		return
	}
}