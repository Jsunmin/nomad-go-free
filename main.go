// http://golang.site/ 참고!
package main

// $GOPATH / $GOROOT 기반으로 모듈을 찾는다! (폴더까지!)
import (
	"fmt"

	"github.com/sunmin/nomad-go-free/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "first value"}
	fmt.Println(dictionary)
	// Search method
	definition, searchErr := dictionary.Search("first")
	if searchErr != nil {
		fmt.Println(searchErr)
	} else {
		fmt.Println(definition)
	}

	// Add method
	addErr := dictionary.Add("second", "second value")
	if addErr != nil {
		fmt.Println(addErr)
	}
	fmt.Println(dictionary)

	// Update method
	updateErr := dictionary.Update("second", "second update value")
	if updateErr != nil {
		fmt.Println(updateErr)
	}
	fmt.Println(dictionary)

	// Delte method
	deleteErr := dictionary.Delete("first")
	if deleteErr != nil {
		fmt.Println(deleteErr)
	}
	fmt.Println(dictionary)
}
