package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(http.ListenAndServe(":3333", nil))
}
