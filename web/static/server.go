package main
import (
	"fmt"
	"net/http"
)

func main() {
	def := http.FileServer(http.Dir("static"))
	http.Handle("/", def)
	http.HandleFunc("/about", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "static/about.html")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8080", nil)
}