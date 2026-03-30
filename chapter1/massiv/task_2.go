package main
import "fmt"

func main() {
	var arr [5][5] int;
	for i := 0; i < 5; i++ {
		if (i%2 == 0) {
			for j := 0; j < 5; j++ {
				arr[i][j] = i*5 + (j+1);
				fmt.Print(arr[i][j], " ");
			}
		} else {
			for k := 4; k >= 0; k-- {
				arr[i][k] = i*5 + (k+1);
				fmt.Print(arr[i][k], " ");
			}
		}
		fmt.Println();
	}
}