// http://golang.site/ 참고!
package main

// $GOPATH / $GOROOT 기반으로 모듈을 찾는다! (폴더까지!)
import (
	"fmt"

	"github.com/sunmin/nomad-go-free/accounts"
)

func main() {
	// 폴더내 file명부터 시작 (파일단위 모듈?!)
	publicAccount := accounts.AccountPublic{Owner: "sunmin", Balance: 1000}
	// 이 경우 프로퍼티 수정이 가능..
	publicAccount.Owner = "ssun?"
	fmt.Println(publicAccount)

	// 함수를 통해 private한 프로터피틀 가진 인스턴스? 생성하는 패턴!
	privateAccount := accounts.NewAccount("sunmin")
	// privateAccount.owner = "ssun?" ~ private 찾을 수 없어!
	// 대신 메서드를 통해 프로퍼티 변경!
	privateAccount.Deposit(500)
	// privateAccount.Withdraw(100)
	err := privateAccount.Withdraw(600) // 원금 초과 에러 발생!
	// 에러 핸들링! ~ 함수에서 받은 에러를 손수 처리해줘야 한다.
	if err != nil {
		// print and exit by os
		// log.Fatalln(err)
		fmt.Println(err)
	}
	// 프로퍼티 리턴하는 메서드 추가
	fmt.Println(privateAccount.Balance(), privateAccount.Owner())
	// 각 프로퍼티 메서드를 혼합해 표현
	fmt.Println(privateAccount) // String 메서드 활용
}
