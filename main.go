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

	user1, err := module.GetUser(1)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(user1)

	todos, err := user1.GetTodosByUserID()
	if err != nil{
		log.Fatalln(err)
	}
	for i, v := range todos{
		fmt.Println(i, v)
	}
	/*
	user1.CreateTodo("this is second todo")

	todo, err := module.GetTodo(2)
	fmt.Println(todo)


	err := u.CreateUser()
	if err != nil{
		log.Fatalln(err)
	}
	



	user1.Password = "22testtest"
	err = user1.UpdateUserPassword()
	*/

}
