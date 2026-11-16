package main 
import "fmt"

type node struct {
	number int
	next *node
}

func main() {
	n1 := node{number: 1};
	n2 := node{number: 2};
	n3 := node{number: 3};
	
	n1.next = &n2;
	n2.next = &n3;

	var current_number *node = &n1;
	for current_number != nil {
		fmt.Println(current_number.number);
		current_number = current_number.next;
	}
}