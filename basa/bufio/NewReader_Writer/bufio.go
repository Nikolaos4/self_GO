package main
import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func main() {
	/*
	rows := []string{"Hello, world", "How are you"}
	file, err := os.Create("bufio.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	writer := bufio.NewWriter(file)
	defer file.Close()

	for _, n := range rows{
		writer.WriteString(n)
		writer.WriteString("\n")
	}
	writer.Flush()
	*/
	file, err := os.Open("bufio.txt")
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
				os.Exit(1)
			}
		}
		fmt.Print(line)
	}
}