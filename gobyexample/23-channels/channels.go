/*
Channels are the pipes that connect concurrent
goroutines. You can send values into channels from one
goroutine and receive those values into another goroutine.
*/
package main

import "fmt"

func main() {
	// Create a new channel with make(chan val-type).
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax.
	// Here we send "ping" to the messages channel we made
	// above, form a new goroutine.
	go func() { messages <- "pings" }()

	// When we run the program the "ping" message is
	// successfully passed from one goroutine to another via our
	// channel.
	msg := <-messages
	fmt.Println(msg)

	// By default sends and receives block until both the sender
	// and receiver are ready. This property allowed us to wait at
	// the end of our program for the "ping" message without
	// having to use any other synchronization.
}

/*
OUTPUT
$ go run channels.go
ping
*/
