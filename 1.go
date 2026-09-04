package main
import (
	"fmt"
)

func main() {
    text := "cwedned"
    pos := 3
    s := "D"
    text_arr := []rune(text)
    help_arr := make([]rune, len(text_arr))
    copy(help_arr, text_arr)
    new_arr := append(help_arr[:pos], []rune(s)...)
    new_arr = append(new_arr, text_arr[pos:]...)
    fmt.Println(string(new_arr))
}
