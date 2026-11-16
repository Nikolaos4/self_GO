package main
import "fmt"

func FirstLetters(s string) string {
	arr := []rune(s)
	new_str := []rune{}
	flag := true
	for i := 0; i < len(arr); i++ {
		if (arr[i] != 32) && (flag == true) {
			new_str = append(new_str, arr[i])
			flag = false
		} else if (arr[i] == 32) {
			flag = true
		}
	}
	return string(new_str)
}

func main() {
	fmt.Println(FirstLetters("Привет мир"))
}