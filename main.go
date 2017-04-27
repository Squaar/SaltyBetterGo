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

	for {
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

		if state.Status == "open" {
			err = client.PlaceBet(1, 10)
			if err != nil {
				panic(err)
			}
			fmt.Println("Bet 10 on Player1.")
		}

		time.Sleep(15 * time.Second)
	}
}
