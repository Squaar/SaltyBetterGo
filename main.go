package main

import (
	"fmt"
	"time"
)

func main() {
	client, err := NewSaltyClient("saltyface@gmail.com", "saltyface")
	if err != nil {
		panic(err)
	}

	balance, err := client.GetWalletBalance()
	if err != nil {
		panic(err)
	}

	state, err := client.GetState()
	if err != nil {
		panic(err)
	}

	fmt.Println("Wallet Balance: ", balance)
	fmt.Println("State: ", state)

	for {
		newBalance, err := client.GetWalletBalance()
		if err != nil {
			panic(err)
		}

		newState, err := client.GetState()
		if err != nil {
			panic(err)
		}

		fmt.Println("Wallet Balance: ", newBalance)
		fmt.Println("State: ", newState)

		if newState.Status == "open" {
			err = client.PlaceBet(1, 10)
			if err != nil {
				panic(err)
			}
			fmt.Println("Bet 10 on Player1.")
		}

		time.Sleep(15 * time.Second)
	}
}
