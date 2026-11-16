package main
import "fmt"

type TextEditor struct {
	content string
}


func (text *TextEditor) Insert(pos int, s string) {
  	text_arr := []rune(text.content)
	new_arr := make([]rune, 0, len(text_arr)+1)
	new_arr = append(new_arr, text_arr[:pos]...)
	new_arr = append(new_arr, []rune(s)...)
	new_arr = append(new_arr, text_arr[pos:]...)
	text.content =string(new_arr)
}

func(text *TextEditor) Delete(pos, lenght int) {
	if (len(text.content) < pos + lenght) {
		fmt.Println("Error")
		return
	}
	rune_arr := []rune(text.content)
	rune_arr = append(rune_arr[:pos], rune_arr[pos + lenght:]...)
	text.content = string(rune_arr)
}

func(text *TextEditor) Content() string{
	text_arr := []rune(text.content)
	new_arr := []rune{}
	for _, val := range text_arr {
		new_arr = append(new_arr, rune(val))
	}
	return string(new_arr)
}

func main() {
	var str string
	fmt.Scanln(&str)
	t := TextEditor{content: str}
	p_t := &t
	p_t.Insert(2, "D")
	fmt.Println(p_t.Content())
}