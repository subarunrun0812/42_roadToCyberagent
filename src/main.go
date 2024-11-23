package main

import (
	"fmt"
	"net/http"
)

func HelloWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", HelloWord)
	fmt.Println("Server is running on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
