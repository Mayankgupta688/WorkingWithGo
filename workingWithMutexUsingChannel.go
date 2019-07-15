package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func workingWithMutexUsingChannel() {
	runtime.GOMAXPROCS(8)

	mutex := make(chan bool, 1)

	for i:=0; i<10; i++ {
		for j:=0; j<10; j++ {

			// The Go routine are now not run in parralel, but the execution indicates that the parallel behavior is somehow reduced...

			// the next go routine is executed only when the previous goroutine unlocks the mutex..

			mutex <- true

			go func() {
				fmt.Println("i: " + strconv.Itoa(i) + " j: " + strconv.Itoa(j))
				<- mutex
			}()
		}
	}

	fmt.Scanln()
}