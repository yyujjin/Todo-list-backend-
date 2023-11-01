package main

import (
	"todo-list-backend/router"
	// "router"

	"github.com/gin-gonic/gin"
)

//main 은 연결돼있어서 저기선 쓸필요없음.

func main() {
	r := gin.Default()
	
	r.GET("/ping", router.Pong)

	// todos -> todos 슬라이스에 있는 값을 전달
	// ['todo1', 'todo2', 'todo3']
	

	//라우터에있는 익명함수를 다른파일로 분리해주는거
	//다른파일에 함수가 있으면 대문자로 해줘야 쓸수가이씀.
	r.GET("/todos", router.Todos) 
	r.POST("/todos", router.PostTodos) 
	r.DELETE("/todos", router.DeleteTodos) 
	
	r.Run() // listen and serve on 0.0.0.0:8080
}