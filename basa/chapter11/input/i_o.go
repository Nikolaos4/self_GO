package main
import (
	"fmt"
	"os"
	"io"
)

type phone string

func(p phone) Read(b []byte) (int, error){
	counter := 0
	for i := 0; i < len(p); i++ {
		if (p[i] >= '0' && p[i] <= '9') {
			b[counter] = p[i]
			counter++
		}
	}
	return counter, io.EOF
}

func main() {
	file, err := os.Open("hello.txt")
	if (err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
	fmt.Println()
	//---------------------------

	p := phone("+32387237822")
	io.Copy(os.Stdout, p)
	fmt.Println()
}