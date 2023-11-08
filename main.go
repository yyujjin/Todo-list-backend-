package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//main 은 연결돼있어서 저기선 쓸필요없음.

func main() {
	r := gin.Default()
	
	r.GET("/ping", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// todos -> todos 슬라이스에 있는 값을 전달
	// ['todo1', 'todo2', 'todo3']
	

	
	type todo struct {
		Id int `json:"id"`
		Name string `json:"name"`
	}


	todos := []todo{{1,"todo1"},{2,"todo2"}}

	r.GET("/todos", func (c *gin.Context) {
	
	 
	fmt.Println(todos)
		c.JSON(200, todos)
	}) 
	
	//5. post : 요소 추가 {id: 3, name: 'todo3'} => 언니가 원하는 형태
	
	r.POST("/todos", func (c *gin.Context) {
		//append로 넣어줄때는 자료형(구조체)을 기입해야 컴퓨터가 알아먹음  ???
			
		var newTodo todo
		//&이거 왜 쓰는지 모르겠고
		//newTodo변수는 그럼 처음엔 그냥 선언됐다가
		//바디에 담긴 값을 변수에 저장하면 계속 값이 바뀌는 건가?
		if err := c.BindJSON(&newTodo); err != nil {
			return
		}

		todos = append(todos,newTodo)
		c.JSON(http.StatusCreated, newTodo )
	}) 
	 
//:id를 클라이언트가 입력하면 string으로 값이 들어오는건가?????
	r.GET("/todos/:id", func (c *gin.Context) {
		//밑에 코드를 하는 이유는 ?'
		//for range로 변수에 담긴 값고 비교해야 하는데
		//todo구조체의 id필드는 string 이니까  ==>>int인데??
		//클라이언트가 요청한 id값 int랑 자료형이 맞지않아 비교가 안돼서  ==>클라이언트가 요청한 값이 string아님??
		//todo 필드 id값을 int로 바꿔주는 코드.  ==>>새로만든 id의 자료형이 string이니까 그걸 int로 바꿔주는 거 아님???
		id,_ := strconv.Atoi (c.Param("id"))  //자료형이 string 이라서  "" 이 안에 쓴건가???
		for _, a := range todos {
			if a.Id == id {
				c.IndentedJSON(http.StatusOK, a)
				return  // 근데 이 return 한 값이 어디가는거임?
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	})

	//슬라이스 중간ㄴ 삭제 법
	// slice = append(slice[:1], slice[2:]...)


	r.DELETE("/todos/:id", func (c *gin.Context) {
		id,_ := strconv.Atoi (c.Param("id"))  
		for index, a := range todos {
			if a.Id == id {
				c.IndentedJSON(http.StatusOK, a)
				todos = append(todos[:index],todos[index+1:]...)
				// new :=todos[index:]
				// fmt.Println(new)
				return  
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	})


	


	r.Run() // listen and serve on 0.0.0.0:8080
}