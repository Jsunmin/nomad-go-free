package accounts

import (
	"errors"
	"fmt"
)

var errNoMoney = errors.New("Can't withdraw")

// 자료구조 및 프로퍼티까지 export 하려면 public setting 필요
// == 대문자화!!!

// AccountPublic struct
type AccountPublic struct {
	// uppercase start = to make public
	Owner   string
	Balance int
}

// AccountPrivate struct
type AccountPrivate struct {
	// not uppercase start = to make private
	owner   string
	balance int
}

// NewAccount creates AccountPrivate
func NewAccount(owner string) *AccountPrivate {
	account := AccountPrivate{owner: owner, balance: 0}
	// 함수에서 생성한 struct를 그대로 넘긴다!
	// = 포인터를 통한 리턴!
	return &account
}

// Deposit method for AccountPrivate
func (apv *AccountPrivate) Deposit(amount int) {
	// 호출하는 인스턴스 자체가 아닌 복사본!
	// fmt.Println(apv, "?")
	apv.balance += amount
	// fmt.Println(apv, "!") 분명히 증가하지만, 원본에는 변화가 없다!

	// -> 원본을 받아오기 위해 apv *AccountPrivate 필요!
}

// Balance method for AccountPrivate
func (apv AccountPrivate) Balance() int {
	return apv.balance
}

// Withdraw method for AccountPrivate
func (apv *AccountPrivate) Withdraw(amount int) error {
	// error handling
	// 더 많은 금액을 인출한다면?
	if apv.balance < amount {
		return errNoMoney
	}
	apv.balance -= amount
	// null / none 같은 역할
	return nil
}

// ChangeOwner method for AccountPrivate
func (apv *AccountPrivate) ChangeOwner(owner string) {
	apv.owner = owner
}

// Owner method for AccountPrivate
func (apv *AccountPrivate) Owner() string {
	return apv.owner
}

// struct로 만든 obj를 프린트할때 지동으로 실행되는 Go 자체적인 메서드
// 그렇기 떄문에 comment가 필요 없다!
// &{sunmin 500} -> sunmin's account.\nHas: 500

func (apv *AccountPrivate) String() string {
	return fmt.Sprint(apv.owner, "'s account.\nHas: ", apv.balance)
}
