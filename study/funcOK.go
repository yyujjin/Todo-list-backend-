package main

import "fmt"

// 파라미터로 받은 두 수를 더한 후 값을 리턴하는 함수
func plus(num1,num2 int) {
	fmt.Println(num1 + num2)
	// return num1+num2
}


func main() {

	plus(1,2)
	
// 익명함수를 만들어서 변수에 담고, 그 변수를 호출
//????????????????????????????????????

//함수를 변수에 담으면 함수가 됨.
//익명함수는 뒤에 바로 호출할수도있고
//변수에 담아서 나중에 호출할수도 있음.  
	result := func (num1,num2 int) {
	fmt.Println(num1+num2)
}
	result(1,2)


// 	type user struct {
// 		name string
// 		age int
// 	}
	
// 	yujin := user{
// 		name: "yujin",
// 		age: 27,
// 	}
// 	fmt.Println(yujin)

// 	// 과목 이름과 점수를 가지는 구조체를 만들기
// 	// 수학, 50

	type subject struct {
		name string
		score int
	}

	// test := subject{
	// 	name : "수학",
	// 	score: 50,
	// }
	// fmt.Println(test)

	test :=subject {"수학",50}
	fmt.Println(test)

}

//
const user = {
	name: 'yujin',
	age: 27
} 
//밑에처럼 그냥 객체만들어쓰는거랑 구조체 만들어서 쓰는거랑 무슨 차이?
//map은 value를 여러개의 자료형을 가질 수 없음.
//하지만 구조체는 한 객체안에 여라 자료형의 value 를 가질 수 있다. 
user := map[string]int{
	"name": 0,
	"age": 27
}

type user struct {
	name string
	age int
}

