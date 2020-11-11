// http://golang.site/ 참고!
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

	// array ~ 크기제한 존재
	namesArray := [5]string{"sun1", "sun2", "sun3"}
	fmt.Println(namesArray)
	namesArray[2] = "sun33"
	namesArray[3] = "sun4"
	fmt.Println(namesArray)
	namesArray[4] = "sun5"
	fmt.Println(namesArray)
	// names[5] = "sun6" ~ 배열 사이즈를 넘어가면서 컴파일 에러

	// 2 slice ~ 크기제한 없는 배열 -> 동적인 사이즈 증가가 가능함
	namesSlice := []string{"sun1", "sun2", "sun3"}
	// slice에 원소를 추가할 때는 append로 원소 추가한 새배열 만든 후 대체함
	namesSlice[2] = "sun33"
	// * 배열의 경우 append활용 X
	namesSlice = append(namesSlice, "sun4")
	fmt.Println(namesSlice)

	// 3 map ~ key/value가 존재하는 자료형 (2곳의 타입이 선정의)
	// 키에 대응하는 값을 신속히 찾는 해시테이블(Hash table)을 구현한 자료구조
	// 아래의 경우 key string, value string
	namesMap := map[string]string{
		"name": "sun1",
		"age":  "12", // 12이면 컴파일 에러 not string..
	}
	fmt.Println(namesMap)
	// iterate 가능!
	for key, value := range namesMap {
		fmt.Println(key, value)
	}

	// 4 struct ~ Custom Data Type을 표현하는데 사용 ~ 자유롭게 value 설정해서 활용
	// 필드들의 집합체이며 필드들의 컨테이너이다. (필드 데이타만 존재, 메서드 존재X)
	// 구조체 선언
	type person struct {
		name    string
		age     int
		favFood []string
	}
	favFood := []string{"chicken", "beer"}
	// 순서대로 value만 쓰던 key&value쌍을 맞춰주는 형식으로 생성
	namesStruct := person{name: "sunmin", age: 28, favFood: favFood}
	fmt.Println(namesStruct)
}
