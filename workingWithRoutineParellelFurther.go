package main
import "runtime"

func main() {

	// If the parellel processing is removed.. it will not execute other Golang routines..

	runtime.GOMAXPROCS(8)

	// Here these processes are Parellel. so they do not wait for main function to pause or wait

	go firstRoutine()
	go secondRoutine()
	go thirdRoutine()
	go firstRoutine()
	go secondRoutine()
	go thirdRoutine()

	println("Go Routine Executed..")
}

func firstRoutine() {
	println("This is First GoRoutine")
}

func secondRoutine() {
	println("This is Second GoRoutine")
}

func thirdRoutine() {
	println("This is Third GoRoutine")
}