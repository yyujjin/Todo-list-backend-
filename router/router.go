package router

import "github.com/gin-gonic/gin"

// pong 일반 함수로 변경
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

var todos = []string { // TODO전역변수로 만들기
	"todo1",
	"todo2", 
	"todo3",
}

func Todos(c *gin.Context) {
	

	c.JSON(200, gin.H{
		"todos": todos, 
	})
}

func PostTodos(c *gin.Context) {
	todos = append (todos,"todo4")
	c.JSON(200, gin.H{
		
		"todos": todos, 

	})
}

func DeleteTodos(c *gin.Context) {
	todos = todos[1:]
	c.JSON(200, gin.H{
		
		"todos": todos, 

	})
}