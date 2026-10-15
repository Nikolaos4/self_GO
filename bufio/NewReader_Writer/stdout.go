package main
import (
	"os"
	"io"
)

type phone string

func(p phone) Read(b []byte) (int, error) {
	count := 0
	for i := 0; i < len(p); i++ {
		if (p[i] >= '0' && p[i] <= '9') {
			b[count] = p[i]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	p := phone("+7-985-207-49-12")
	io.Copy(os.Stdout, p)
}