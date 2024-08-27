package router

import(
	"net/http"
	"main/handlers"
)

func InitRouters(){

	http.HandleFunc("/api/v1/resource",handlers.GetTodos)
	//↑ http://8000/api/v1/resource?user_id=1 みたいにアクセスすると handlers.GetTodosが実行される
}