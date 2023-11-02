package main

import "fmt"

// 1. person struct 생성
// name과 age를 가짐
// a := 10

type person struct {
	name string
	age int
}

func main() {
	// 구조체 생성
var persons = person{"yujin",28}

	// 2 함수 호출
	value(persons)
	fmt.Println(persons)
	// 3 함수 호출
	refer(&persons)
	fmt.Println(persons)
}

// 2. pass by value -> age를 1 더해주는 함수. 파라미터로 구조체를 받아야함
func value(a person) {
a.age += 1
fmt.Println(a.age)
}
// 3. pass by refer -> age에다가 1더하고 마찬가지로 구조체 받음
func refer (a *person) {
	a.age += 1
	fmt.Println(a.age)
}