package main
import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	/*
	io.Copy(os.Stdout, resp.Body)
	*/

	for true {
		bs := make([]byte, 1024)
		n, err := resp.Body.Read(bs)
		if err != nil {break}
		fmt.Println(string(bs[:n]))
	}	
}