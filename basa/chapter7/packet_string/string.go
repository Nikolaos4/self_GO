package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World"
	//изменение регистра
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))

	//поиск префикса/суфикса строки
	fmt.Println(strings.HasPrefix(str, "He"))
	fmt.Println(strings.HasSuffix(str, "ld"))

	//поиск текста
	fmt.Println(strings.Contains(str, "World"))

	new_str := "Hello World. Good bye World"
	//количество вхождений в строке
	fmt.Println(strings.Count(new_str, "World"))

	//преобразование массива в строку
	slice1 := []string{"Привет", "прекрасный", "мир"}
	str_from_slice1 := strings.Join(slice1, " ")
	fmt.Println(str_from_slice1)

	//замена подстроки
	replace_str := strings.Replace(str, "World", "work", 1)
	fmt.Println(replace_str)

	//преобразование строки в массив
	telefon_str := "+7-985-207-49-12"
	array := strings.Split(telefon_str, "-")
	fmt.Println(array)
}
