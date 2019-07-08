package main
import "time"
import "runtime"

func main() {

	// This will allow the process to utilize all the cores of the System..

	runtime.GOMAXPROCS(8)

	// Now each Go Lang routine can be executed over different core..

	go alphabetGenerator()
	println("Go Routine Executed..")
}

func alphabetGenerator() {
	for alpha := byte('a'); alpha <= byte('z'); alpha++ {
		go println(string(alpha))
	}
}