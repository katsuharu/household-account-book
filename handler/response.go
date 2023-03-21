package handler

// Expense defines model for Expense.
type Expense struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// HTTPError defines model for error.
type HTTPError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
