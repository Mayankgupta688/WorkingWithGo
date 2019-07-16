package main

import "fmt"

func main() {

	// Here we have created a Buffered Chanel, which do not stops as soon as data is recieved.

	firstChannel := make(chan string, 1)
	secondChannel := make(chan string, 1)

	secondChannel <- "Mayank Gupta"

	// Using Select Keyword, we are telling "Golang" to read data from the first channel which is available to deliver the data

	select {
		case dataFromFirstChannel := <- firstChannel: 
			fmt.Println(dataFromFirstChannel)
		case dataFromSecondChannel := <- secondChannel: 
			fmt.Println(dataFromSecondChannel)
	}
}