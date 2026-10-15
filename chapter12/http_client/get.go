package main
import (
	"fmt"
	"net/http"
	"io"
	"os"
)

func main() {
	client := http.Client{}
	resp, err := client.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}