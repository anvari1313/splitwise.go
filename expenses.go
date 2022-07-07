package splitwise

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
)

// Expenses contains method to work with expense resource
type Expenses interface {
	// Expenses returns current user's expenses
	// Expenses(ctx context.Context) ([]Expense, error)

	// // ExpenseByID returns information of an expense identified by id argument
	// ExpenseByID(ctx context.Context, id uint64) (*Expense, error)

	// CreateExpense Creates an expense. You may either split an expense equally (only with group_id provided), or
	// supply a list of shares.
	//If providing a list of shares, each share must include paid_share and owed_share, and must be identified by one
	// of the following:
	//email, first_name, and last_name
	//user_id
	//Note: 200 OK does not indicate a successful response. The operation was successful only if errors is empty.
	CreateExpenseSplitEqually(ctx context.Context, expense ExpenseSplitEqually) ([]Expense, error)
	CreateExpenseByShare(ctx context.Context, expense ExpenseByShare) ([]Expense, error)
}

type Expense struct {
	Cost           string `json:"cost"`
	Description    string `json:"description"`
	Details        string `json:"details"`
	Date           string `json:"date"`
	RepeatInterval string `json:"repeat_interval"`
	CurrencyCode   string `json:"currency_code"`
	CategoryId     uint32 `json:"category_id"`
	GroupId        uint32 `json:"group_id"`
}

type ExpenseSplitEqually struct {
	Expense
	SplitEqually bool `json:"split_equally"`
}

type ExpenseByShare struct {
	Expense
	ByShare map[string]interface{}
}

type createExpenseResponse struct {
	Expenses []Expense `json:"expenses"`
}

func (c client) CreateExpenseSplitEqually(ctx context.Context, expense ExpenseSplitEqually) ([]Expense, error) {

	url := c.baseURL + "/api/v3.0/create_expense"

	body, err := json.Marshal(expense)
	if err != nil {
		return nil, err
	}

	token, err := c.AuthProvider.Auth()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var response createExpenseResponse
	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return response.Expenses, nil
}

func (c client) CreateExpenseByShare(ctx context.Context, expense ExpenseByShare) ([]Expense, error) {

	url := c.baseURL + "/api/v3.0/create_expense"

	err := checkByShareValues(expense.ByShare)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(expense)
	if err != nil {
		return nil, err
	}

	token, err := c.AuthProvider.Auth()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var response createExpenseResponse
	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return response.Expenses, nil
}

func checkByShareValues(items map[string]interface{}) error {

	regex := "(users__(0__(user_id$|paid_share$|owed_share$)))|(users__[0-9]__(paid_share$|owed_share$|first_name$|last_name$|email$))"
	reg, err := regexp.Compile(regex)

	if err != nil {
		return errors.New("wrong regex")
	}

	finalMatch := true
	for k := range items {
		match := reg.MatchString(k)
		finalMatch = match && finalMatch
	}
	if finalMatch {
		return nil
	}
	return errors.New("wrong payload")

}
