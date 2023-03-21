package repository

import (
	"context"

	"github.com/katsuharu/household-account-book/domain/object/expense"
)

type Expense interface {
	Create(ctx context.Context, expense *expense.Expense) (*expense.Expense, error)
}
