package main
import (
	"fmt"
	"net/http"
)

type msg string
func(m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
}

func main() {
	ms := msg("Hello, word!")
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", ms)
}
