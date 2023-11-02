package main

import "fmt"

// 1. grade struct 만들기. score 필드를 가지고, score는 map[string]int 타입임
type grade struct {
	name string
	score1 map[string]int
	score2 map[string]int
	score3 map[string]int
	score4 map[string]int
	score5 map[string]int
	score6 map[string]int
}

func main() {
	// var a = map[string]int{
	// 	"math":  90,
	// 	"music": 100,
	// }

	student1Name := "yujin"
	var student1Score = map[string]int{}
	student1Score["math"]=90
	student1Score["music"]=100
	fmt.Println(student1Name, student1Score)

	// 2. grade 사용해서 변수 만들기
	var a = grade{}
	a.name = "yujin" 
	a.score1=map[string]int{}
	a.score2=map[string]int{}
	a.score3=map[string]int{}
	a.score4=map[string]int{}
	a.score5=map[string]int{}
	a.score6=map[string]int{}
	a.score1["math"]=90
	a.score1["music"]=100
	fmt.Println(a)
	// grade 구조체타입으로 선언 -> map 초기화 -> map에 값 넣기

	b := newStruct()
	b.score["math"] = 90
	b.score["music"] = 100
	fmt.Println(b)

	c := newStruct()
	b.score1["math"] = 70
	// grade 생성자로 선언 -> map 에 값 넣기
}



func newStruct() *grade { 
	d := grade{} 
	d.score1 = map[string]int{}
	d.score2 = map[string]int{}
	d.score3 = map[string]int{}
	d.score4 = map[string]int{}
	d.score5 = map[string]int{}
	return &d 
}