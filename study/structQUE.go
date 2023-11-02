package main

import (
	"fmt"
)

// user 구조체 만들기. age와 name을 필드로 가짐
type user struct {
	age int
	name string
}

func main() {
	// users := user{}
	// users.age = 16
	// users.name = "yujin"

	// users := user{16, "yujin"}
	//변수만들어서 넣은거 
	// yujin := user{age : 16, name : "yujin"}
	




	var userList []user // [] 구조체 이름 => 구조체 슬라이스 만들기
	// 구조체를 요소로 가지는 슬라이스를 만들겠다는 뜻. ! 

	//변수안만들고 바로 할수도있음.
	//변수 만들고 하는 방법도 있음? 
	//이게 변수 만들어서 배열에 추가하는 거 
	tempUser := user{age : 16, name : "yujin"}
	tempUser2 := user{age : 16, name : "yujin"}
	userList = append(userList, tempUser, tempUser2)
	//이게 변수 안만들고 그냥 넣는 거 (이게 더 효율 적)
	userList = append(userList, user{age : 16, name : "yujin"}, user{age : 16, name : "yujin"})
		fmt.Println(userList)


		//구조체 슬라이스 만듣는 이유 
		//{age : 16, name : "yujin"} => 이건 그냥 객체 하나만 있는 거 
		//배열안에 요소 하나안에 객체하나가 들어가있다. 
		//요소의 자료형이 객체 
		//한 배열안에 여러객체들을 관리하고 싶어서
		// user1  := {age : 16, name : "yujin"}
		// user2 := {age : 16, name : "yujin"}
		// user3 := {age : 16, name : "yujin"}

		// //ddddd(코드 여러줄...)


		// user4 := {age : 16, name : "yujin"}

		// userList := [{age : 16, name : "yujin"},{age : 16, name : "yujin"},{age : 16, name : "yujin"},{age : 16, name : "yujin"}]
		//[{age : 16, name : "yujin"},{age : 16, name : "yujin"},{age : 16, name : "yujin"}]
}