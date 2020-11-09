// 패키지 선언 필수
package main

import (
	"fmt"
)

// go entry function
func main() {
	// 기본적으로 export 되는 모듈은 대문자로 시작함 (Println)
	// 소문자는 private한 것으로 취급
	fmt.Println("Hello world!")
	// Go = type언어!, 변수 옆에 타입을 지정해야 함
	const name string = "sunmin"
	// name = "ssun" ~ const 수정 불가
	/*
	** 함수내 & var(변수) 기준으로 쓸 수 있는 변수 초기화 축약형
	** name1 := "sunmin"
	** = var name1 string = "sunmin"
	** 타입은 초기화 값을 통해 자동으로 찾아줌 ( 이 경우 string )
	**
	 */
	name1 := "sunmin"
	name1 = "ssun" // true 같은 다른 타입 변경은 불가!
	fmt.Println(name)
	fmt.Println(name1)
}
