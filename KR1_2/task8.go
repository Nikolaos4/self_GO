package main
import "fmt"

func len_of_word(word string) int{
	return len(word);
}

func main() {
	words := []string{"cat", "dog", "elephant", "fox", "giraffe"};
	var max int = 0;
	var max_word string;
	for i := 0; i < len(words); i++ {
		if len_of_word(words[i]) > max {
			max = len_of_word(words[i]);
			max_word = words[i];
		}
	}
	fmt.Print(max_word);
}