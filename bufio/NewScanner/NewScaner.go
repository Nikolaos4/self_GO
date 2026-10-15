package main
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	arr := []int{}
	for scanner.Scan() {
		word := scanner.Text()
		if (word == "0") {
			break
		}
		num, err := strconv.Atoi(word)
		if err != nil {
			fmt.Println(err)
		}
		arr = append(arr, num)
	}
	fmt.Println(arr)
}