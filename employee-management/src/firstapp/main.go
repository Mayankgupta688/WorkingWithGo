package main

import "net/http"

func main() {

	// The request is pointed out by Address, so that it can be accessed directly
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8000", nil)
}