package main

import (
	"fmt"
)

//RunServer starts the connections with the Discord API using discordgo and gorilla/websocket
func RunServer() {
	fmt.Println("* Logging in!")
	Session.Token, err = Session.Login(Server.Username, Server.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Open websocket connection
	fmt.Println("* Opening websocket connection!")
	err = Session.Open()
	if err != nil {
		fmt.Println(err)
	}

	// Do websocket handshake.
	fmt.Println("* Do websocket handshake!")
	err = Session.Handshake()
	if err != nil {
		fmt.Println(err)
	}

	// Listen for events.
	fmt.Println("* Listening")
	Session.Listen()
	return
}
