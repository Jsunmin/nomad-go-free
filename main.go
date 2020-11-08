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
}
