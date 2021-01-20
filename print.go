package main

import (
	"fmt"
	"net/http"
)

//IndexHandler writes a message to /
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is so cool :)")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8000", nil)

}
