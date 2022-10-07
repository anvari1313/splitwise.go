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

type ActionBy struct {
	Id                 uint32 `json:"id"`
	FirtsName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	RegistrationStatus string `json:"registration_status"`
	Picture            struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"picture"`
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

type ExpenseResponse struct {
	Expense
	Id                     uint32   `json:"id"`
	FriendshipId           uint32   `json:"friendship_id"`
	Repeats                bool     `json:"repeats"`
	EmailReminder          bool     `json:"email_reminder"`
	EmailReminderInAdvance uint8    `json:"email_reminder_in_advance"`
	NextRepeat             string   `json:"next_repeat"`
	CommentsCount          string   `json:"comments_count"`
	Payment                string   `json:"payment"`
	Transactionconfirmed   string   `json:"transaction_confirmed"`
	CreatedAt              string   `json:"created_at"`
	CreatedBy              ActionBy `json:"created_by"`
	UpdatedAt              string   `json:"updated_at"`
	UpdatedBy              ActionBy `json:"updated_by"`
	DeletedAt              string   `json:"deleted_at"`
	DeletedBy              ActionBy `json:"deleted_by"`
	Repayments             []struct {
		From   uint32 `json:"from"`
		To     uint32 `json:"to"`
		Amount string `json:"amount"`
	} `json:"repayments"`
	Category struct {
		Id   uint32 `json:"id"`
		Name string `json:"Name"`
	} `json:"category"`
	Receipt struct {
		Large    string `json:"large"`
		Original string `json:"original"`
	} `json:"receipt"`
	Users []struct {
		User struct {
			Id        uint32 `json:"id"`
			FirtsName uint32 `json:"first_name"`
			LastName  uint32 `json:"last_name"`
			Picture   uint32 `json:"picture"`
		} `json:"user"`
		UserId     uint32 `json:"user_id"`
		PaidShare  string `json:"paid_share"`
		OwedShare  string `json:"owed_share"`
		NetBalance string `json:"net_balance"`
	} `json:"users"`
	Comments []struct {
		Id           uint32 `json:"id"`
		Content      string `json:"content"`
		CommentType  string `json:"comment_type"`
		RelationType string `json:"relation_type"`
		RelationId   uint32 `json:"relation_id"`
		CreatedAt    string `json:"created_at"`
		DeletedAt    string `json:"deleted_at"`
		User         string `json:"user"`
	} `json:"comments"`
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

func (c client) GetExpenseCurrentUser(ctx context.Context) (*ExpenseResponse, error) {

	url := c.baseURL + "/api/v3.0/get_expenses"

	token, err := c.AuthProvider.Auth()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var response ExpenseResponse
	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
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
