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
	fmt.Println(friends)
}
