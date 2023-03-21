package handler

import (
	"github.com/katsuharu/household-account-book/util/logs"
	"net/http"

	"github.com/katsuharu/household-account-book/usecase"
	"github.com/labstack/echo/v4"
)

type ExpenseHandler interface {
	// CreateExpense (POST /expenses)
	CreateExpense(ctx echo.Context) error
}

type expenseHandlerWrapper struct {
	l *logs.Logger
	u usecase.ExpenseUseCase
}

func NewExpenseHandler(u usecase.ExpenseUseCase, l *logs.Logger) ExpenseHandler {
	return &expenseHandlerWrapper{u: u, l: l}
}

func (w expenseHandlerWrapper) CreateExpense(ctx echo.Context) error {
	req := new(ReqBodyExpense)
	if err := ctx.Bind(req); err != nil {
		w.l.Error("failed to bind request", err)
		return ctx.JSON(http.StatusBadRequest, HTTPError{
			Code:    001,
			Message: "パラメータのバインドに失敗しました。",
		})
	}

	result, err := w.u.Create(ctx.Request().Context(), req.Name)
	if err != nil {
		w.l.Error("failed to create expense", err)
		return ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    002,
			Message: "費目の登録に失敗しました。",
		})
	}

	return ctx.JSON(http.StatusCreated, Expense{
		ID:   result.ID,
		Name: req.Name,
	})
}
