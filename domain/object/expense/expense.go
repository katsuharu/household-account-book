package expense

import (
	"fmt"
	"time"
)

type Expense struct {
	ID        ID
	Name      Name
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewExpense(name string, currentTime time.Time) (*Expense, error) {
	id, err := newID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new id: %w", err)
	}

	expenseName, err := newName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to generate name: %w", err)
	}

	return &Expense{
		ID:        id,
		Name:      expenseName,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}, nil
}
