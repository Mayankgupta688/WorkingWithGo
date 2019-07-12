package main

import "fmt"

func main() {

	ch := make(chan string);

	// Here the channel waits for some data to be recieved

	// But since there is no other channel to give data.. Therefore it goes in the deadlock

	// All threads waiting for something to happen in the future, but nothing would happen since no other goroutine is there..

	fmt.Println(<-ch)
}