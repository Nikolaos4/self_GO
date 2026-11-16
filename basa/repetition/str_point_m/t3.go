package main
import "fmt"

type MyString struct {
	s string
}

func Reverse(ms MyString) string{
	arr := []rune(ms.s)
	reverse_arr := []rune{}
	for i := len(arr)-1; i >= 0; i-- {
		reverse_arr = append(reverse_arr, arr[i])
	}
	return string(reverse_arr)
}

func main() {
	ms := MyString{s: "Привет мир"}
	fmt.Println(Reverse(ms))
}