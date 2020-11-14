package mydict

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("Not Found")
var errAlreadyExists = errors.New("Already Exists")
var errCantUpdate = errors.New("Can't Update Non-existing word")
var errCantDelete = errors.New("Can't Delete Non-existing word")

//Dictionary type
type Dictionary map[string]string // 이렇게 타입에 alias를 지정할 수 있다!

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	// go map 내장 메서드 ( 값과 존재여부 리턴 )
	value, exits := d[word]
	if exits {
		return value, nil
	}
	return "", errNotFound

	// dictionary["hello"] = "hello" 기능이지만,
	// 에러 핸들링 등 alias한 특정 타입에 대한 메서드로 단단한 코드 작성 가능!
}

// Add for a word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err { // JS와 달리 waterfall 이 없다! 한 case에서 끝
	case errNotFound:
		d[word] = def
		fmt.Println("case1?")
	case nil:
		fmt.Println("case2?")
		return errAlreadyExists
	}
	return nil

	// 리시버에서 *를 붙여주지 않음 (struct는 붙여줬는데..)
	// Nico:  we don't need the * on the receiver because maps on Go are automatically using *
}

// Update for a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete for a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
