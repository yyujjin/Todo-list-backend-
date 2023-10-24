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
//todo구조체를 여러개 사용하기위해 [배열 ] 로 관리
var todos = []todo{
	{id: 1, name: "todo1"},
	{id: 2, name: "todo2"},
}


func Todos(c *gin.Context) {
	
	fmt.Print(todos)
	c.JSON(200, gin.H{
		"todos": todos, 
	})
}

func PostTodos(c *gin.Context) {
	// 1. todo3을 슬라이스에 추가
	todos = append(todos,todo{3,"todo3"})
	// 2. todos 슬라이스를 출력(응답)
	fmt.Println(todos)

	c.JSON(200, gin.H{
		
		"todos": todos, // todo1~todo3

	})
}

func DeleteTodos(c *gin.Context) {
	todos = todos[1:]
	c.JSON(200, gin.H{
		
		"todos": todos, 

	})
}