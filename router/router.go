package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// pong 일반 함수로 변경
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}



// [
// 
// ]

type todo struct {
	id int
	name string
}

type todoList struct {
	todos []todo // 자료형 : 배열이라는 뜻? 근데 뒤에 todo는 왜 붙은거임 
}

// 1. todos를 리턴하는 메소드
func (tl todoList) returnTodos() []todo{
	return tl.todos
}
// 2. todos에 {id: 3, name: "todo3"} 추가하는 메소드
func (tl todoList) addTodos(){
	tl.todos = append(tl.todos,todo{id:3,name: "todo3"})
}

//todo구조체를 여러개 사용하기위해 [배열 ] 로 관리
// var todos = []todo{
// 	{id: 1, name: "todo1"},
// 	{id: 2, name: "todo2"},
// }

var todos = todoList{[]todo{
	{123,"박유진"},
	{456,"박수현"},
}}


func Todos(c *gin.Context) {
	// 
	t := todos.returnTodos()
	fmt.Println(t)
	todos.addTodos()
	fmt.Println(todos.returnTodos())
	c.JSON(200, gin.H{
		"todos": t, 
	})
}

func PostTodos(c *gin.Context) {
	// 1. todo3을 슬라이스에 추가
	//todos = append(todos,todo{3,"todo3"})
	// 2. todos 슬라이스를 출력(응답)
	fmt.Println(todos)

	c.JSON(200, gin.H{
		
		//"todos": todos, // todo1~todo3

	})
}

func DeleteTodos(c *gin.Context) {
	//todos = todos[1:]
	c.JSON(200, gin.H{
		
		//"todos": todos, 

	})
}