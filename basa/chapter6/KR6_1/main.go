package main

import (
	"probability/count_package"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Количество человек: ")
	reader := bufio.NewReader(os.Stdin)
	var count_of_people int
	fmt.Scanln(&count_of_people)
	
	var hash_names = map[string]int {}
	var name string
	var age int
	for i:= 0; i < count_of_people; i++ {
		fmt.Println("Введите данные: ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		
		_, err := fmt.Sscan(line, &name, &age)
		if (err != nil) {
			fmt.Println("Error!")
		}
		hash_names[name] = age
		
	}
	
	fmt.Print("\nВывод всех данных: \n")
	for key, value := range hash_names {
		fmt.Printf("name: %s, age: %d \n", key, value)
	}

	current_pair := check(hash_names)
	fmt.Printf("\nКоличество подходящих пар: %d \n", current_pair)

	answer := count_package.Probability(current_pair, len(hash_names))
	fmt.Printf("Ответ: %d", answer)
}