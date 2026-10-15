package main
import (
	"fmt"
	"io"
)

type phone string

type Phone struct{}

func(p phone) Read(b []byte) (int, error){
	count := 0
	for i := 0; i < len(p); i++ {
		if (p[i] >= '0' && p[i] <= '9') {
			b[count] = p[i]
			count++
		}
	}
	return count, io.EOF
}

func(p Phone) Writer(b []byte) (int, error){
	count := 0
	for i := 0; i < len(b); i++ {
		if (b[i] >= '0' && b[i] <= '9') {
			fmt.Printf("%c ", b[i])
		}
	}
	fmt.Println()
	return count, io.EOF
}

func main() {
	p := phone("8-800-555-35-35")
	arr := make([]byte, len(p))
	p.Read(arr)
	fmt.Println(string(arr))

	//----------------------

	writer := Phone{}
	writer.Writer(arr)
}