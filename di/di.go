package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	// fmt.Printf("hello, %s", name)
	fmt.Fprintf(writer, "hello, %s", name)
}

func MyGreetHanlder(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "Chris")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHanlder))
}
