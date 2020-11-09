// 패키지 선언 필수
package main

import (
	"fmt"
	"strings"
)

// go entry function
func main() {
	fmt.Println(multiply(2, 2))
	// 사용없이 단순 선언만 하면, 컴파일러가 에러로 잡아줌
	totalLength, upperName := lenAndUpper("sunmin")
	// 만약에 한 값만 쓰고 싶다면? : _로 처리! (convention?!)
	// totalLength, _ := lenAndUpper("sunmin")
	fmt.Println(totalLength, upperName)
	repeatMe("s", "su", "sun", "sunm", "sunmi", "sunmin")
	// naked return 함수 (함수 실행 -> length, uppercase 생성 -> defer 실행 -> printLn 실행)
	fmt.Println("nakedAndDeferLenAndUpper")
	fmt.Println(nakedAndDeferLenAndUpper("sunmin"))
	deferTest()
}

// 함수 정의시 in/output에 대한 타입 정의가 필수다!
func multiply(a int, b int) int {
	return a * b
}

// 1개 이상의 값 (multiple vale)을 리턴할 수 있다!
func lenAndUpper(name string) (int, string) {
	// 모두 go의 표준 라이브러리로 존재함!
	return len(name), strings.ToUpper(name)
}

// 여러 args를 어떻게 받을까? -> ...type 으로 정의!
func repeatMe(words ...string) {
	// 모두 go의 표준 라이브러리로 존재함!
	fmt.Println(words)
}

// naked return 리턴 이름까지 전부 타입화! (func 선언시 length와 uppercase로 리턴 값을 지정)
func nakedAndDeferLenAndUpper(name string) (length int, uppercase string) {
	// 리턴 후 실행되는 코드!
	defer fmt.Println("I'm done")
	// 변수는 위에서 초기화되었으니 바로 사용!
	length = len(name)
	uppercase = strings.ToUpper(name)
	// 리턴 명시 필요 X (이미 함수 선언시 지정함)
	return // length, uppercase (해도 작동함)
}

// defer
func deferTest() {
	// 함수가 전부 실행되고 defer가 돈다!
	defer fmt.Println("in defer~")
	fmt.Println("no retrun func@")
}
