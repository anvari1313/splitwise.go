package splitwise

import (
	"context"
)

// Expenses contains method to work with expense resource
type Expenses interface {
	// Expenses returns current user's expenses
	Expenses(ctx context.Context) ([]Expense, error)

	// ExpenseByID returns information of an expense identified by id argument
	ExpenseByID(ctx context.Context, id uint64) (*Expense, error)

	// CreateExpense Creates an expense. You may either split an expense equally (only with group_id provided), or
	// supply a list of shares.
	//If providing a list of shares, each share must include paid_share and owed_share, and must be identified by one
	// of the following:
	//email, first_name, and last_name
	//user_id
	//Note: 200 OK does not indicate a successful response. The operation was successful only if errors is empty.
	CreateExpense(ctx context.Context, dto *CreateCommentDTO) ([]Expense, error)
}

type Expense struct {
}

type CreateExpenseDTO struct {
}
