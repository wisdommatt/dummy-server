package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Welcome to cloudnotte !"))
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
