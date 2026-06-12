package main

import "fmt"

func main() {
	var value rune = 0xd09f
	fmt.Printf("%c \n", value)
    fmt.Printf("%d \n", value)
    fmt.Println(value)

	//перебор строки, преобразованной в срез рун
	str := "Привет мир"
	rune_slice := []rune(str) //преобразования строки в срез рун
	for i := 0; i < len(rune_slice); i++ {
		fmt.Printf("%c", rune_slice[i])
	}

	//Преобразование рун в строку
	myrune := [] rune{0x41f, 0x440, 0x438, 0x432, 0x435, 0x442, 0x020, 0x43c, 0x438, 0x440}
    new_str := string(myrune)
	fmt.Println(new_str)
}