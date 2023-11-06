package main

func main() {
	
	//new를 썼을 땐 객체에 값 넣는 법 . 으로만 됨. 
	// c := new(nums)
	// c.a = 1
	// c.b = 2
	// result2 := addNumStruct(c)
	// println(result2)
	// c1 := nums{1,2}
	c2 := nums{a:1,b:2}
	addNumStruct(&c2)
	fmt.Println(c1,c2)
}

// *적으면 포인터로만 사용 됨
// 그냥 함수로는 안됨. 할경우 따로 생성해야 함. 
//* 자체가 포인터  
// 포인터 함수.
func addNumStruct(s *nums) int { 
	// nums{1,2} 다른 메모리 공간에 똑같은 값이 복사됨.
	return s.a + s.b
}


	
//& = 메모리 주소
// & 이 없으면 그 메모리 주소에 담긴 값을 전달
// & 이 있으면 메로지 주소 자체를 전달

//* = 메모리 주소안에 있는 값을 사용하겠다라는 의미

//a,b 메모리 주소를 매개변수로 넘겼어
//그 이후에 *a *b 쓰는 이유는 매개변수에 담긴 메모리를 꺼내써야하기떄문에 
//*int = int가 담긴 메모리 주소를 받겠다. =>그저 자료형 
	// a := 1
	// b := 2
	// result := addNum(&a, &b) // & 메모리 주소
	// println(result)
func addNum(a,b *int) int {
	return *a + *b
}