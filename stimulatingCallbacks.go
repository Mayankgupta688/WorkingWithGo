package main

import (
	"fmt"
)

func main() {
	callbackChannel := make(chan *UserDetails)
	firstUser := new(UserDetails)
	go saveUserDetails(firstUser, callbackChannel)
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