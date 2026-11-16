package protocol

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

//Информацию о количестве уровней надо заносить в файл
//Записывать сообщение лога в соответсвующий файл

var level_map = map[string] int {
	"info": 0,
	"warn": 0,
	"error": 0,
}


func check_file(filename string) int{
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer f.Close()
	var count int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		count++
	}
	return count
}

func level_check(level string, part_arr []string) string {

	level = strings.ToLower(level)
	var flag string
	if _, ok := level_map[level]; ok {
		message := strings.Join(part_arr, " ")
		name_file := level + ".txt"
		file, err := os.OpenFile(name_file, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
			flag = "Error while open file"
			return flag
		}
		defer file.Close()
		fmt.Fprintln(file, message)
		flag = fmt.Sprintf("The message was added with log %s", level)
	} else {
		flag = "Error: the level is not found"
	}
	return flag
}

func get_statistic() string{
	info_level := check_file("info.txt")
	warn_level := check_file("warn.txt")
	error_level := check_file("error.txt")
	result := fmt.Sprintf("level info: %d; level warn: %d; level error: %d", info_level, warn_level, error_level)
	return result
}

func ProtocolCheck(arg_from_str []string) string{
	if len(arg_from_str) == 0 {
		return "Error: empty command"
	}
	protocol_user := arg_from_str[0]
	protocol_user = strings.ToLower(protocol_user)
	var result string
	switch(protocol_user) {
		case "log":
			if len(arg_from_str) < 3 {
				result = "Error: usage: LOG <level> <message>"
				return result
			}
			result = level_check(arg_from_str[1], arg_from_str[2:])
		case "stats":
			result = get_statistic()
		case "quit":
			result = "-1"
		default:
			result = "The protocol is not recognized"
	}
	return result
}