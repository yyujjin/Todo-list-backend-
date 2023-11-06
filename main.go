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
	

	// 1. 클라이언트가 /todos를 호출하면 접근하면 함수가 실행됨.
	// 2.  이 라우터는 배열을 내보내는 라우터.
	// [{id: 1, name: 'todo1'}] => 원하는 형태
	//필드 네임을 앞글자 대문자로 해야지 화면에 나타나는데
	//보통은 소문자로 시작하기때문에 옆에 json을 넣어줘서 소문자로 만들어줘. 
	type todo struct {
		Id int `json:"id"`
		Name string `json:"name"`
	}

	//3.todos라는 배열을 만드는데 구조체를 자료형으로가진 배열
	//4.todo구조체만 넣을수있는 구조라서 []todo{todo{1,"todo"}}=>이렇게 안해줘도 됨. 
	todos := []todo{{1,"todo1"},{2,"todo2"}}

	r.GET("/todos", func (c *gin.Context) {
	
	//JSON = 응답해주는 함수
	//todos 배열을 내보내겠다는 의미 
	fmt.Println(todos)
		c.JSON(200, todos)
	}) 
	
	//5. post : 요소 추가 {id: 3, name: 'todo3'} => 언니가 원하는 형태
	
	r.POST("/todos", func (c *gin.Context) {
		//append로 넣어줄때는 자료형(구조체)을 기입해야 컴퓨터가 알아먹음
		//POST로 요청하면 어펜드로 값을 넣어주고
		//JOSN은 추가된 todos배열 자체를 보여준다는 뜻.
		//newtodo라는 변수를 만들고 데이터타입은 todo구조체를 한다
		var newTodo todo

		//클라이언트가 바디로 보낸값을 변수에다 넣어주는 코드 밑
		if err := c.BindJSON(&newTodo); err != nil {
			return
		}
		// 그리고 값이 담긴 그 변수를 todos 배열에 넣어줌.
		todos = append(todos,newTodo)
		//클라이언트한테 니가준거 잘들어갔다고 응답해줌
		c.JSON(http.StatusCreated, newTodo )
	}) 

	r.GET("/todos/:id", func (c *gin.Context) {
		//밑에 코드를 하는 이유는 ?'
		//for range로 변수에 담긴 값고 비교해야 하는데
		//todo구조체의 id필드는 string 이니까 
		//클라이언트가 요청한 id값 int랑 자료형이 맞지않아 비교가 안돼서
		//todo 필드 id값을 int로 바꿔주는 코드.
		id,_ := strconv.Atoi (c.Param("id"))  
		for _, a := range todos {
			if a.Id == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	})

	

	r.Run() // listen and serve on 0.0.0.0:8080
}