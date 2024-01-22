package models

import "time"

type Expense struct {
	ExpenseId   uint ` gorm:"primary_key"`
	Title       string
	Amount      float64
	ExpenseFile *string
	Description string
	CreatedBy   int
	ApprovedBy  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      int
}

func (Expense) TableName() string {
	return "expense_requests"
}
