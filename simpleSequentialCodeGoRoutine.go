package main

import "time"

func simpleSequentialCodeGoRoutine() {
	go func() {
		println("Hello")
	}()

	go func() {
		println("World")
	}()

	time.Sleep(1000 * time.Millisecond)
}