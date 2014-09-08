package main

import (
	"fmt"
)

func main() {
	client, err := AuthClient("mdumford99@gmail.com", "hamsters99")
	if err != nil {
		panic(err)
	}

	fmt.Println(WalletBalance(client))
}
