package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(8)
	
	ch := make(chan string)
	doneChannel := make(chan string)

	go alphabetGenerator(ch)
	go alphabetReader(ch, doneChannel)
	<- doneChannel
}


func alphabetGenerator(ch chan string) {
	for alpha := byte('a'); alpha <= byte('z'); alpha++ {
		ch <- string(alpha)
	}

	close(ch)
	
}


func alphabetReader(ch chan string, doneChannel chan string) {
	for l := range ch {
		println("Data")
		println(l)
	}
	
	doneChannel <- "Hello"
}