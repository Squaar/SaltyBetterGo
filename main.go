package main

import (
	"fmt"
	// "github.com/murz/go-socket.io"
	// "net/http"
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

	if state.Status == "open" {
		err = PlaceBet(1, 200, client)
		if err != nil {
			panic(err)
		}
		fmt.Println("Bet placed")
	}

	// An attempt at creating a socket.io client...
	// I don't know how to open a conection to the saltybet server
	// I think the messages it sends are empty anyway. Can just check status periodically instead.

	// sio := socketio.NewServer(nil)
	// var msg socketio.Message
	// http.Handle("/socket.io/", http.StripPrefix("/socket.io",
	// 	sio.Handler(func(conn *socketio.Conn) {
	// 		for {
	// 			if err = conn.Receive(&msg); err != nil {
	// 				panic(err)
	// 			}
	// 			switch msg.Type() {
	// 			case socketio.MessageJSON, socketio.MessageText:
	// 				fmt.Println("normal message: %s", msg.String())
	// 			case socketio.MessageEvent:
	// 				fmt.Println("event message: %s", msg.String())
	// 			}
	// 		}
	// 	})))

	// if err = http.ListenAndServe("www-cdn-twitch.saltybet.com:8000", nil); err != nil { // BROKEN
	// 	panic(err)
	// }
}
