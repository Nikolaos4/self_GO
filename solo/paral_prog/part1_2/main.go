package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Result struct {
	Filename string
	count    int
	err      error
}

func countLines(filename string, ch chan<- Result) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		ch <- Result{filename, 0, err}
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		str := scanner.Text()
		str = strings.TrimSpace(str)
		if str != "" {
			count++
		}
	}
	ch <- Result{filename, count, nil}
}

func Close(ch chan Result, wg *sync.WaitGroup) {
	wg.Wait()
	close(ch)
}

func main() {
	var wg sync.WaitGroup
	var ch chan Result = make(chan Result)

	wg.Add(1)
	go func() {
		defer wg.Done()
		countLines("error.txt", ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		countLines("info.txt", ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		countLines("warn.txt", ch)
	}()

	go Close(ch, &wg)

	for r := range ch {
		fmt.Printf("File: %s, count: %d, error: %v\n", r.Filename, r.count, r.err)
	}

	//
	arr_filename := []string{"error.txt", "info.txt", "warn.txt"}
	count := 0
	for line := range filter(readLines(gen(arr_filename)), "тест") {
		fmt.Println(line)
		count++
	}
	fmt.Println(count)
}

func gen(name_slice []string) <-chan string {
	ch_name := make(chan string)
	go func() {
		defer close(ch_name)
		for _, v := range name_slice {
			ch_name <- v
		}
	}()
	return ch_name
}

func readLines(ch_name <-chan string) <-chan string {
	ch_str_of_file := make(chan string)
	go func() {
		defer close(ch_str_of_file)
		for name := range ch_name {
			helpRead(name, ch_str_of_file)
		}
	}()
	return ch_str_of_file
}

func helpRead(name string, ch_str_of_file chan<- string) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		ch_str_of_file <- str
	}
}

func filter(lines <-chan string, word string) <-chan string {
	ch_good_str := make(chan string)
	go func() {
		defer close(ch_good_str)
		for str := range lines {
			arr_word := strings.Fields(str)
			for _, v := range arr_word {
				if v == word {
					ch_good_str <- str
					break
				}
			}
		}
	}()
	return ch_good_str
}
