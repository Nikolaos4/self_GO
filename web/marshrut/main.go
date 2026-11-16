package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/about", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "information")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "hello.html")
	})

	fmt.Println("Server is listening")
	http.ListenAndServe("localhost:8080", nil)
}


/*
type httpHandle struct {
	message string
}

func(hh httpHandle) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, hh.message)
}

func main() {
	h0 := httpHandle{"Вечер"}
	h1 := httpHandle{"Урааааа"}
	h2 := httpHandle{"Гооол!"}

	http.Handle("/", h0)
	http.Handle("/what", h1)
	http.Handle("/yes", h2)

	fmt.Println("Server is listening")
	http.ListenAndServe("localhost:8080", nil)
}	
*/