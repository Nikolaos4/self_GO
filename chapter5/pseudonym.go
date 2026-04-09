//псевдоним просто определяет новое название для типа
package main 
import "fmt"

type mile uint 
type function func([3]int) int 

func get_sum(arr [3]int, sum_arr function) int {
	return sum_arr(arr);
}

func sum(arr [3]int) int{
	summa := 0;
	for _, value := range arr {
		summa += value;
	}
	return summa;
}

func main() {
	var distance mile = 5;
	distance += 3;
	fmt.Println(distance);

	fmt.Println(get_sum([3]int{2,1,4}, sum));
}