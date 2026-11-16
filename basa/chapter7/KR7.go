package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

//проверка наличия ключевых строк
func presence(str string) bool {
    return strings.Contains(str, "word")
}

//считает количество 
func count(str string) int {
    return strings.Count(str, "word")
}

//заменяет слово на другое
func replace(str string) string{
    return strings.ReplaceAll(str, "word", "work")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

    slice_from_string := strings.Split(line, " ")
    for index := range slice_from_string {
        slice_from_string[index] = strings.ToLower(slice_from_string[index])
    }

    new_str := strings.Join(slice_from_string, " ")
    fmt.Println(new_str)

    flag_presence := presence(new_str)
    if flag_presence && count(new_str) > 1 {
        answer := replace(new_str)
        fmt.Printf("Ответ: %s", answer)
    }
}