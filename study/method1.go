package main

import "fmt"

// 1. width, height 필드를 가지는 Triangle 구조체 생성
type Triangle struct {
	width, height float32
}

// method!
func (t Triangle) getArea()float32 {
	return t.width * t.height / 2
}

// 2. width * height / 2 결과를 반환하는 getArea 함수 구현 (메소드 아님!)
func getAreaNotMethod (a *Triangle) float32 {
	return a.width * a.height / 2
}
// 3. main에서 구조체 사용해서 변수만들고 함수 호출 후 결과 출력
func main() {

	tri1 := Triangle {15,5}
	tri2 := getAreaNotMethod(&tri1)
	fmt.Println(tri2)
	
	tri3 := tri1.getArea()
	fmt.Println(tri3)

}