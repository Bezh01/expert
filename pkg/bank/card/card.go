package card

import (
	"bank/pkg/bank/types"
)

const bonusLimit = 5_000_00

// IssueCard func
func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}

	return card
}

// AddBonus func
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYears int) {
	bonus := ((int(card.MinBalance) * percent / 100 * daysInMonth) / daysInYears)
	if !card.Active {
		return
	}
	if card.Balance < 0 {
		return
	}
	if bonus > bonusLimit {
		return
	}
	card.Balance += types.Money(bonus)
	return
}

//Deposit func
func Deposit(card *types.Card, amount types.Money) {
	if card.Active && amount <= 50_000_00 && amount > 0 {
		card.Balance += amount
		return
	}
	return
}

//Withdraw func
func Withdraw(card types.Card, amount types.Money) types.Card {
	if !card.Active {
		return card
	}
	if card.Balance < amount {
		return card
	}
	if amount < 0 {

		return card
	}
	if amount > 20_000_00 {
		return card
	}
	card.Balance = card.Balance - amount
	return card
}

// Total func
func Total(cards []types.Card) types.Money {
	var sum types.Money = types.Money(0)
	for _, card := range cards {
		sum += card.Balance
	}
	return sum
}

// PaymentSource func
func PaymentSources(cards []types.Card) []types.PaymentSource {
	
	payment:= []types.PaymentSource{}	
	for _, bezh := range cards {
	if bezh.Balance>0 && bezh.Active {

		payment= append(payment, types.PaymentSource{Number:string(bezh.PAN), Balance:bezh.Balance})

	}	
}	
    return payment
}
    



