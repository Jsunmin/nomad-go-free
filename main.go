// 패키지 선언 필수
package main

import (
	"fmt"
)

func main() {
	// low level programming
	a := 2
	b := a
	// memory address check : &var
	fmt.Println(&a, &b, a, b)
	a = 10
	fmt.Println(&a, &b, a, b)
	// c = a's memory address
	c := &a
	// 각각 c메모리 주소, c저장값, c저장값(메모리주소)의 저장값
	fmt.Println(&c, c, *c)
	/*
	** c는 a의 메모리주소를 참조하기 때문에 primaryValue 변경도 함께 적용됨
	** 1. 이를 활용해서, 메모리 저장공간을 아낄 수 있다!!
	** 2. 반대로 c를 통해 a값을 변경시킬 수도 있다!!
	 */
	a = 5
	fmt.Println(&a, &b, a, b)
	fmt.Println(&c, c, *c)
	*c = 9
	fmt.Println(&a, &b, a, b)
	fmt.Println(&c, c, *c)
}
