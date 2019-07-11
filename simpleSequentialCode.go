package main

import "time"

func main() {
	func() {
		println("Hello")
	}()

	func() {
		println("World")
	}()

	time.Sleep(1000 * time.Millisecond)
}