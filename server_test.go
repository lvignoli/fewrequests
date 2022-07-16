package fewrequests_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lvignoli/fewrequests"
)

func ExampleListenAndServeN() {
	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello,")
	})

	worldHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "world.")
	})

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/world", worldHandler)

	log.Fatal(fewrequests.ListenAndServeN(4, ":8080", nil))
}
