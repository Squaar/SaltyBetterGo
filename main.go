package main

import (
	"fmt"
	"time"
)

func main() {
	client, err := AuthClient("saltyface@gmail.com", "saltyface")
	if err != nil {
		panic(err)
	}

	balance, err := WalletBalance(client)
	if err != nil {
		panic(err)
	}

	state, err := GetState(client)
	if err != nil {
		panic(err)
	}

	fmt.Println("Wallet Balance: ", balance)
	fmt.Println("State: ", state)

	for {
		newBalance, err := WalletBalance(client)
		if err != nil {
			panic(err)
		}

		newState, err := GetState(client)
		if err != nil {
			panic(err)
		}

		fmt.Println("Wallet Balance: ", newBalance)
		fmt.Println("State: ", newState)

		if newState.Status == "open" {
			err = PlaceBet(1, 10, client)
			if err != nil {
				panic(err)
			}
			fmt.Println("Bet 10 on Player1.")
		}

		time.Sleep(15 * time.Second)
	}
}
