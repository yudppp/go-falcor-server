package main

import (
	"fmt"
	"net/http"

	"falcor"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/model.json", falcor.FalcorHandler)
	http.ListenAndServe(":9090", nil)
}
