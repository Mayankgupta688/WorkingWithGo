package main

import "fmt"

func creatingBasicChannelWithOtherError() {

	ch := make(chan string);

	// Sending a message in the channel, causes a channel to block untill someone recieves it. 

	// Since the main channel is blocked on this point untill anyone access it, therefore the next line is not executed.

	// When a gorouting sends a message in the channel, it waits for someone to drain that message

	// Channel has no capacity to store messages, so it need to drained first before proceeding further.

	ch <- "Hello"

	fmt.Println(<-ch)
}