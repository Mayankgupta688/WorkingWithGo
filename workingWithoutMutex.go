package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func workingWithoutMutex() {
	runtime.GOMAXPROCS(8)

	for i:=0; i<10; i++ {
		for j:=0; j<10; j++ {
			go func() {
				fmt.Println("i: " + strconv.Itoa(i) + " j: " + strconv.Itoa(j))
			}()
		}
	}

	fmt.Scanln()
}