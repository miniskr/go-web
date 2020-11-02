package main

import "net/http"

type myHandleer struct {}

func (m *myHandleer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello world"))
}

func main() {

	mh := myHandleer{}

	server := http.Server{
		Addr: "localhost:8888",
		Handler: &mh,	
	}

	server.ListenAndServe()
}