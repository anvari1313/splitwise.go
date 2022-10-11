package splitwise

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// Expenses contains method to work with expense resource
type Expenses interface {
	// Expenses returns current user's expenses
	Expenses(ctx context.Context) ([]ExpenseResponse, error)

	// ExpenseByID returns info about an expense choose by id
	ExpenseByID(ctx context.Context, id uint64) (ExpenseResponse, error)

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
	FirstName          string `json:"first_name"`
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
	PaidUserID uint64 `json:"users__0__user_id"`
	OwedUserID uint64 `json:"users__1__user_id"`
	PaidShare  string `json:"users__0__paid_share"`
	OwedShare  string `json:"users__1__owed_share"`
}

type ExpenseResponse struct {
	Expense
	ID                     uint64   `json:"id"`
	FriendshipID           uint64   `json:"friendship_id"`
	Repeats                bool     `json:"repeats"`
	EmailReminder          bool     `json:"email_reminder"`
	EmailReminderInAdvance int8     `json:"email_reminder_in_advance"`
	NextRepeat             string   `json:"next_repeat"`
	CommentsCount          uint     `json:"comments_count"`
	Payment                bool     `json:"payment"`
	TransactionConfirmed   bool     `json:"transaction_confirmed"`
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
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Picture   struct {
				Medium string `json:"medium"`
			} `json:"picture"`
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

type expensesResponse struct {
	Expenses []ExpenseResponse `json:"expenses"`
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

	err = c.checkError(res)
	if err != nil {
		return nil, err
	}

	var response createExpenseResponse
	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return response.Expenses, nil
}

func (c client) CreateExpenseByShare(ctx context.Context, expense ExpenseByShare) ([]Expense, error) {
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

	err = c.checkError(res)
	if err != nil {
		return nil, err
	}

	var response createExpenseResponse
	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return response.Expenses, nil
}

func (c client) Expenses(ctx context.Context) ([]ExpenseResponse, error) {
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

	err = c.checkError(res)
	if err != nil {
		return nil, err
	}

	var response expensesResponse
	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return response.Expenses, nil
}

type expenseByIDResponse struct {
	Expense ExpenseResponse `json:"expense"`
}

func (c client) ExpenseByID(ctx context.Context, id uint64) (ExpenseResponse, error) {
	url := c.baseURL + "/api/v3.0/get_expense/" + strconv.FormatUint(id, 10)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return ExpenseResponse{}, err
	}

	token, err := c.AuthProvider.Auth()
	if err != nil {
		return ExpenseResponse{}, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, err := c.client.Do(req)
	if err != nil {
		return ExpenseResponse{}, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	err = c.checkError(res)
	if err != nil {
		return ExpenseResponse{}, err
	}

	var response expenseByIDResponse
	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return ExpenseResponse{}, err
	}

	return response.Expense, nil
}
