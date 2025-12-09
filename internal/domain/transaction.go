package domain

import "time"

type TransactionType string

const (
	TransactionTypeIncome  TransactionType = "income"
	TransactionTypeExpense TransactionType = "expense"
)

type Transaction struct {
	ID          int64
	Date        time.Time
	Description string
	Amount      float64
	Category    string
	Type        TransactionType
}

type CategoryTotal struct {
	Category string
	Total    float64
}
