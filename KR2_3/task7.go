package main
import "fmt"

func main() {
	month := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
	time_of_year := [4]string{"Зима","Весна","Лето","Осень"};
	var num int;
	fmt.Scan(&num);

	if (1 <= num && num <= 2 || num == 12) {
		fmt.Printf("Время года: %s, количество дней: %d", time_of_year[0], month[num-1]);
	} else if (3 <= num && num <= 5) {
		fmt.Printf("Время года: %s, количество дней: %d", time_of_year[1], month[num-1]);
	} else if (6 <= num && num <= 8) {
		fmt.Printf("Время года: %s, количество дней: %d", time_of_year[2], month[num-1]);
	} else if (9 <= num && num <= 11) {
		fmt.Printf("Время года: %s, количество дней: %d", time_of_year[3], month[num-1]);
	}
	
}