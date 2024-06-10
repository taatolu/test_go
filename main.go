package main

import (
	"fmt"
	"log"
	"main/module"
	_"main/config"
)


func main (){
	log.Println("this is a Test")
	fmt.Println(module.Db)

	u := &module.User{}
	u.Name = "Yusaku"
	u.Email = "test@example.com"
	u.Password = "testtest"

	u.CreateUser()
}
