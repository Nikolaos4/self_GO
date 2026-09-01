package main
import (
	"bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Введите числа через пробел: ")
    input, _ := reader.ReadString('\n')          
    input = strings.TrimSpace(input)
	arr := strings.Split(input, " ")
	fmt.Printf("%T", arr)
}