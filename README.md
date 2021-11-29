# Splitwise Golang SDK

![Test](https://github.com/anvari1313/splitwise.go/actions/workflows/test.yml/badge.svg)

A community driven Golang SDK for [Splitwise](https://splitwise.com) 3rd-party APIs.

## How to use it?

1. You should get the package via ```go get``` command:

~~~
go get -u github.com/anvari1313/splitwise.go
~~~

2. Register your application [here](https://secure.splitwise.com/apps) and obtain an API key for your app. 

3. Put your API key in your code:
~~~go
package main

import (
  "context"
  "fmt"

  "github.com/anvari1313/splitwise.go"
)

func main() {
  auth := splitwise.NewAPIKeyAuth("PUT_YOUR_API_KEY_HERE")
  client := splitwise.NewClient(auth)
	
  currentUser, err := client.CurrentUser(context.Background())
  if err != nil {
	  panic(err)
  }
  
  fmt.Println(currentUser)
}
~~~
