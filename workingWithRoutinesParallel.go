package main
import "runtime"

func workingWithRoutinesParallel() {

	// This will allow the process to utilize all the cores of the System..

	runtime.GOMAXPROCS(8)

	// Now each Go Lang routine can be executed over different core..

	go alphabetGeneratorParallel()
	println("Go Routine Executed..")
}

func alphabetGeneratorParallel() {
	for alpha := byte('a'); alpha <= byte('z'); alpha++ {
		go println(string(alpha))
	}
}