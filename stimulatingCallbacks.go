package main

import (
	"fmt"
)

func main() {
	callbackChannel := make(chan *UserDetails)
	firstUser := new(UserDetails)
	go saveUserDetails(firstUser, callbackChannel)

	// Here the channel will pause the main goroutine untill the data is not available in the channel.

	// Once the data is recieved, channel will resume the execution again
	
	newUser := <- callbackChannel
	fmt.Println(newUser.userName)
}

type UserDetails struct {
	userName string
	userAge int
}

func saveUserDetails(userData *UserDetails, callback chan *UserDetails) {
	userData.userName = "Mayank"
	userData.userAge = 30
	callback <- userData
}