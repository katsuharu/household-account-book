package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/katsuharu/household-account-book/domain/object/expense"
	"github.com/katsuharu/household-account-book/domain/repository"
)

type ExpenseUseCase interface {
	Create(ctx context.Context, name string) (*ExpenseDTO, error)
}

type useCase struct {
	r repository.Expense
}

type ExpenseDTO struct {
	ID   string
	Name string
}

func NewExpenseUseCase(r repository.Expense) ExpenseUseCase {
	return &useCase{
		r: r,
	}
}

// Create 費目を登録
func (u useCase) Create(ctx context.Context, name string) (*ExpenseDTO, error) {
	e, err := expense.NewExpense(name, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to generate expense: %w", err)
	}

	result, err := u.r.Create(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("failed to create expense: %w", err)
	}

	return &ExpenseDTO{
		ID:   string(result.ID),
		Name: result.Name.String(),
	}, nil
}
