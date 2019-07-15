package main

import "fmt"

func creatingBasicChannelWithoutError() {


	// Now this channel is created with a buffer.. which can hold 2 items..

	ch := make(chan string, 2);

	ch <- "Hello"

	ch <- "World"

	fmt.Println(<-ch)

	fmt.Println(<-ch)
}