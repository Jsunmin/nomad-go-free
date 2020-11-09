// 패키지 선언 필수
package main

import (
	"fmt"
)

// go entry function
func main() {
	result1 := canIDrinkIf(18)
	fmt.Println(result1)
	result2 := canIDrinkSwitch(18)
	fmt.Println(result2)
}

// IF
func canIDrinkIf(age int) bool {
	// if 문 형식 + Go에서는 if block내 변수 선언이 가능하다!
	if varInIf := age + 2; varInIf > 18 {
		fmt.Println(varInIf, "true?")
		return true
	}
	return false
}

// SWITCH
func canIDrinkSwitch(age int) bool {
	// switch 문 형식 + Go에서는 switch block내 변수 선언이 가능하다!
	// switch varInIf2 := age + 2; varInIf2 {
	// case 10:
	// 	return false
	// case 18:
	// 	fmt.Println(varInIf2, "switch here!")
	// 	return true
	// case 20:
	// 	fmt.Println(varInIf2, "switch here?")
	// 	return false
	// }
	switch {
	case age+2 < 10:
		return false
	case age+2 == 18:
		return true
	case age+2 >= 20:
		return false
	}
	return false
}
