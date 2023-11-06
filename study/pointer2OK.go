package main

import "fmt"

type nums struct {
	a, b int
}


// 2개의 차이점 : 객체를 어디서 선언할꺼냐의 차이
// 2개중 아무거나 써도 상관없지만 생성자를 썼을때 편한부분이 있음 
func newNums() *nums {
	c2 := nums{a: 1, b: 2}
	return &c2
}

func newNumsNotPointer() nums {
	c2 := nums{a: 1, b: 2}
	return c2
	//c2는 그저 변수이름 
	//변수 자체를 내보낸다는게 아니라 변수에 담긴 값을 내보낸다는 의미 
}

func main() {
	c1 := nums{}
	c1.a = 1
	c1.b = 2
	fmt.Println(addNum2(&c1))
//c2에는 return 값이 들어감 
//c2는 뉴넘스에서 만들어진 애를 가져온거니까 가짜야.
	c2 := newNums()
	//이미 c2 자체가 포인터기때문에 주소를 한번더 넘길필요없음 &c2할 필요 없음. 
	// addNum2(&c2)
	addNum2(c2)

	//c1이랑 c3에 들어가는 값이 똑같애 근데 생성자를 안쓰면 번거롭다.

	// c3 := nums{}
	// c3.a = 1
	// c3.b = 2

	c3 := newNumsNotPointer()
	addNum2(&c3)



}
//return을 해서 밖에서 fmt.Print()이걸 할꺼냐 
//아니면 안에서 print써서 밖에서는 그냥 함수만 적어도 호출되게 할꺼냐 
func addNum2(s *nums) int {
	return s.a + s.b
	// fmt.Print(s.a + s.b)
}