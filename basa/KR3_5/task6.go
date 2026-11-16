package main
import "fmt"

func get_rone(words string) rune{
	var first rune;
	for _, v := range words {
		first = v;
		break;
	}
	return first;
}

func groupByFirstLetter(slice []string) map[rune][]string{
	result_cart := map[rune][]string {};
	for _, v := range slice {
		first_letter := get_rone(v); 
		//if _, ok := result_cart[first_letter]; ok == false {}
		result_cart[first_letter] = append(result_cart[first_letter], v);
	}
	return result_cart;
}

func output(cart map[rune][]string) {
	for key, value := range cart {
		fmt.Printf("%q: %s\n", string(key), value);
	}
}

func main() {
	slice_of_words := []string {"apple", "banana", "apricot", "blueberry", "cherry"};
	result := groupByFirstLetter(slice_of_words);
	output(result);
}
