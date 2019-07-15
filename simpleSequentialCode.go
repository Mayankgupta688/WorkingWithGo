package main

import "time"

func simpleSequentialCode() {
	func() {
		println("Hello")
	}()

	func() {
		println("World")
	}()

	time.Sleep(1000 * time.Millisecond)
}