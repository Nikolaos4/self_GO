package main

import "fmt"

func main() {
	str := "Привет, мир"
	//длина строки
	l := len(str)
	fmt.Println(l) //20, так как русские символы кадируются 2 байтами

    //перебор строки через индексы первого байта рун и числовой код символа(rune). В этом примере я числовой код преобразовал сразу в символ
    for b, value := range str {
        fmt.Printf("%d \t %c \n", b, value)
    }

	//перебор строки по байтам
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	//преобразование срезов бацтов строку
	binary_str := []byte {208, 159, 209, 128, 208, 184, 208, 178, 208, 181, 209, 130, 32, 208, 188, 208, 184, 209, 128}
	new_string := string(binary_str)
	fmt.Println(new_string)
}