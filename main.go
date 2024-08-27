package main

import (
	"net/http"
	"log"
	_"main/module"
	_"main/config"
	"main/router"
)


func main (){
	
	router.InitRouters()
	log.Fatal(http.ListenAndServe(":8000", nil))
	
}
