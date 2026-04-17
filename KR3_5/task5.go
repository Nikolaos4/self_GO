package main
import "fmt"

func rotateRight(slice *[]int, number int) {
	if len(*slice) > number {
		*slice = append((*slice)[len(*slice) - number:], (*slice)[:len(*slice) - number]...);
	} else {
		k := number%len(*slice);
		*slice = append((*slice)[len(*slice) - k:], (*slice)[:len(*slice) - k]...);
	}
}


func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8};
	rotateRight(&slice, 16);
	fmt.Println(slice);
}