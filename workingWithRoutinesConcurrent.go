package main
import "time"
import "runtime" 

func main() {
	go alphabetGenerator()

	// While the main function is kept to sleep, another Go Routine executes and then the function returns ..
	
	time.Sleep(1000 * time.Millisecond)
	println("Go Routine Executed..")
}

func alphabetGenerator() {
	for alpha := byte('a'); alpha <= byte('z'); alpha++ {
		go println(string(alpha))
	}
}