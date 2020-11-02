package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello world"))
	})

	http.ListenAndServe("localhost:8888", nil)
}