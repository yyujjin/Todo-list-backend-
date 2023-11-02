package main

import "fmt"

func main() {
	var a int = 10
	fmt.Println(a) // 10
	changeNumber(&a) // &: 주소
	fmt.Println(a) // 1
	notChangeNumber(a)
	fmt.Println(a) //1
}

func changeNumber(num *int) { // *: 직접 참조
	*num = 1
}

func notChangeNumber(num int) { // 1. 값이 복사됨(메모리에 새로운 공간 생김) 2. 참조가 복사됨 (원래 있던 공간 그대로 사용)
	num = 1
}