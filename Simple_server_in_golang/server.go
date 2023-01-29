package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", SimpleServer)
	http.ListenAndServe(":3000", nil)
}

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yen guru samachara %s", r.URL.Path[1:])
}