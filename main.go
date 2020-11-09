// 패키지 선언 필수
package main

import (
	"fmt"
)

// go entry function
func main() {
	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
}

//
func superAdd(numbers ...int) int {
	total := 0
	// 배열 Loop 접근 ( index, value 제공함)
	// 1, range 활용
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	// 2, for문 활용
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
	return total
}
