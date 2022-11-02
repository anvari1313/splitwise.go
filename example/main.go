package main

import (
	"context"
	"fmt"
	"os"

	"github.com/anvari1313/splitwise.go"
)

func main() {
	auth := splitwise.NewAPIKeyAuth(os.Getenv("API_KEY"))
	client := splitwise.NewClient(auth)

	userExamples(client)
	groupExamples(client)
	friendsExamples(client)
	expensesExamples(client)
	otherExamples(client)
}

func userExamples(client splitwise.Client) {
	currentUser, err := client.CurrentUser(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(currentUser)

	userByID, err := client.UserByID(context.Background(), currentUser.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userByID)

	updatedUser, err := client.UpdateUser(context.Background(), currentUser.ID, splitwise.UserFirstNameField("Ahmad"), splitwise.UserLastNameField("Anvari"))
	if err != nil {
		panic(err)
	}
	fmt.Println(updatedUser)
}

func groupExamples(client splitwise.Client) {
	groups, err := client.Groups(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(groups)

	group, err := client.GroupByID(context.Background(), 110)
	if err != nil {
		panic(err)
	}
	fmt.Println(group)
}

func friendsExamples(client splitwise.Client) {
	friends, err := client.Friends(context.Background())
	if err != nil {
		panic(err)
	}

	for _, friend := range friends {
		fmt.Println(friend.ID, friend.FirstName, friend.LastName, friend.Groups)
	}

	success, err := client.DeleteFriend(context.Background(), 123)
	if err != nil {
		panic(err)
	}

	fmt.Println("Delete fiend:", success)
}

func expensesExamples(client splitwise.Client) {
	expensesRes, err := client.Expenses(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", expensesRes)

	if len(expensesRes) != 0 {
		exp, err := client.ExpenseByID(context.Background(), expensesRes[0].ID)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", exp)
	}

	userShares := []splitwise.UserShare{
		{
			UserID:    27163610,
			PaidShare: "15000.00",
			OwedShare: "7500.00",
		},
		{
			UserID:    58839462,
			PaidShare: "0.00",
			OwedShare: "7500.00",
		},
	}

	expenses, err := client.CreateExpenseByShare(
		context.Background(),
		splitwise.Expense{
			Cost:         "15000.00",
			Description:  "کافه امروز عصر",
			CurrencyCode: "IRR",
			GroupId:      0,
		},
		userShares,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", expenses)
}

func otherExamples(client splitwise.Client) {
	currencies, err := client.Currencies(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", currencies)

	categories, err := client.Categories(context.Background())
	if err != nil {
		panic(err)
	}

	categoriesRecursivePrint(0, categories)
}

func categoriesRecursivePrint(level int, categories []splitwise.Category) {
	for _, category := range categories {
		for i := 0; i < level; i++ {
			fmt.Print("\t")
		}

		if len(category.Subcategories) == 0 {
			fmt.Printf("%d - %s\n", category.ID, category.Name)
		} else {
			fmt.Printf("%d - %s:\n", category.ID, category.Name)
			categoriesRecursivePrint(level+1, category.Subcategories)
		}
	}
}
