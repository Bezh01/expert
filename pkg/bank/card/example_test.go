package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
}
func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 5_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 500000
}
func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}
func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 22_000_00, Active: true}, 21_000_00)
	fmt.Println(result.Balance)
	// Output: 2200000
}

func ExampleDeposit_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 3000000
}

func ExampleDeposit_inactive() {
	card := types.Card{Balance: 0, Active: false}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 0
}

func ExampleDeposit_limit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 52_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus_positive() {
	card := types.Card{Balance: 10_000_00, MinBalance: 10_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1002465
}

func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 10_000_00, MinBalance: 10_000_00, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleAddBonus_negative() {
	card := types.Card{Balance: -10_000_00, MinBalance: 0, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: -1000000
}

func ExampleAddBonus_limit() {
	card := types.Card{Balance: 10_000_00, MinBalance: 10000000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: 1024657
}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}

	fmt.Println(Total(cards))

	//Output: 2000000
}
