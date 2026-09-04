package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	rows := []string{"Hello, word", "How are you"}
	file, err := os.Create("rows.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	writer := bufio.NewWriter(file)
	defer file.Close()

	for _, row := range rows {
		writer.WriteString(row)
		writer.WriteString("\n")
	}

	writer.Flush()
}