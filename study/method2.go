package main

import (
	"fmt"
)

// 1. student 구조체는 이름과 수학 성적, 국어, 영어, 음악 성적을 필드로 가지고 있음
type student struct {
	name string
	math, language, english, music int 
}
// 2. 학생 이름을 리턴하는 메소드 구현
func (a student) returnName() string {
	return a.name
}
// 3. 학생 평균 성적을 리턴하는 메소드 구현
func (a student) returnGrade() int {
	return (a.math + a.language + a.english + a.music) / 4 
}
// 평균 성적을 리턴하는 함수 구현!
func returnGrade (math,language,english,music int) int {
	return (math + language + english + music) / 4 
}
func main() {
	student1 := student{"yujin",20,50,40,90}
	// student1.
	fmt.Println(student1.returnGrade())
	fmt.Println(student1.returnName())

	// returnGrade()
	student1.returnGrade() // 유지보수가 편하다
	result := returnGrade(student1.language,student1.math,student1.english,student1.music)
	fmt.Println(result)
}