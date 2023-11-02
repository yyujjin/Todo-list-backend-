package main

import "fmt"
func main () {
	num := 10
	// 1번 함수 호출
	value(num)
	fmt.Println(num) // 10
	// 2번 함수 호출
	reference(&num)
	fmt.Println(num) // 11

}

// 두 함수 다 num에다가 1을 더함
// 1. pass by value
func value (num int) {
	num += 1
	fmt.Println(num)
}

// 2. pass by reference
func reference (num *int) {
	*num += 1
	fmt.Println(*num)
}