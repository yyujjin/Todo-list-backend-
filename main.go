package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	
	
	type todo struct {
		Id int `json:"id"`
		Name string `json:"name"`
	}


	todos := []todo{{1,"todo1"},{2,"todo2"}}

	r.GET("/todos", func (c *gin.Context) {
	
	 
	fmt.Println(todos)
		c.JSON(200, todos)
	}) 
	
	
	r.POST("/todos", func (c *gin.Context) {
			
		var newTodo todo
		//포인터 함수란 말 자체가 없다.
		// 그냥 함수에 매개변수가 포인터다?
		//newTodo변수는 그럼 처음엔 그냥 선언됐다가
		//바디에 담긴 값을 변수에 저장하면 계속 값이 바뀌는 건가?
		//변수는 함수가 실행될때만 살아있음. => 함수끝나면 newtodo 사라짐 
		
		//바디에 값이 잘 담겼으면 err = nil
		//바디에 값이 안담겼으면 err = [!= nil] 이 상태가 된다. 
		// 함수 만들때 이 사람이 그냥 이렇게 만들어 놨다.
		if err := c.BindJSON(&newTodo); err != nil {
			return
		}

		//이게 풀어 쓴거 
		// err := c.BindJSON(&newTodo);
		// if err != nil {
		// 	return // 함수안에서 retrun 은 함수 종료 
		// }

 
		todos = append(todos,newTodo)
		c.JSON(http.StatusCreated, newTodo )
	}) 
	 
//:id를 클라이언트가 입력하면 string으로 값이 들어오는건가?????
	r.GET("/todos/:id", func (c *gin.Context) {
	//get라우터를 만든사람이 :id 는 string으로 들어오도록 만들어 놨다.
		
		id,_ := strconv.Atoi (c.Param("id"))  //자료형이 string 이라서  "" 이 안에 쓴건가??? 키 값은 "" 안에 넣어주는게 법칙
		for _, a := range todos {
			if a.Id == id {
				c.IndentedJSON(http.StatusOK, a)
				return  
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