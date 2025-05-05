package transaction

type Transaction struct {
	From   string
	To     string
	Amount float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - t.Amount
		}
		if t.To == name {
			return currentBalance + t.Amount
		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0.0)
}

func Reduce[A, B any](collection []A, f func(balance B, t A) B, initialValue B) B {
	var result = initialValue
	for _, val := range collection {
		result = f(result, val)
	}
	return result
}
