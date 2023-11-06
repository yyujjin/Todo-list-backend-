package main

import (
	"fmt"
)

// student 구조체 생성
type student struct {
	name string
	grade1 map[string]int
	grade2 map[string]int
	grade3 map[string]int
}
// name: string, grade1: map[string]int, grade2: map[string]int, grade3: map[string]int


func main() {
	// elly student 생성.
	
ellystudent := student{}
ellystudent.name = "elly"
//map초기화 해줘야함.
ellystudent.grade1 = map[string]int{}
ellystudent.grade1["math"]=50
ellystudent.grade2 = map[string]int{}
ellystudent.grade2["math"]=30
ellystudent.grade3 = map[string]int{}
ellystudent.grade3["math"]=60

fmt.Println(ellystudent)

yujinstudent :=student{}
yujinstudent.name = "yujin"
yujinstudent.grade1 = map[string]int{}
yujinstudent.grade1["math"]=10
yujinstudent.grade2 = map[string]int{}
yujinstudent.grade2["math"]=20
yujinstudent.grade3 = map[string]int{}
yujinstudent.grade3["math"]=30

fmt.Println(yujinstudent)
	// name, grade1~3에 값 넣기

// ??????????????????????????????????
	// 여기엔 왜 * & 이게 없음. ? 
	//없어도 됨. 
	//대신 
sook := newStruct2()
sook.name = "sook"
sook.grade1["math"]=60
sook.grade2["math"]=70
sook.grade3["math"]=80
fmt.Println(sook)
}

// 생성자 함수 만들기 
//그냥 값을 내보내는 생성자라서 포인터안씀
//지금은 그냥 생성자 함수를 만들때 값을 내보내는 생성자를 만들수도 있고 포인터를 내보낼수있는 생성자를 만들수도 있구나!
//라는 걸 알면 됨. 
func newStruct2() student {
	var a = student{}
	a.grade1 = map[string]int{}
	a.grade2 = map[string]int{}
	a.grade3 = map[string]int{}
	return a
}