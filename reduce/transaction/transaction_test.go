package transaction

import "testing"

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From:   "Chris",
			To:     "Riya",
			Amount: 100,
		},
		{
			From:   "Adil",
			To:     "Chris",
			Amount: 25,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
	AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
}
