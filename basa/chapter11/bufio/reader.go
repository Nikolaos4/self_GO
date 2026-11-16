package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	file, err := os.Open("rows.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
			}
		}
		fmt.Print(line)
	}
}