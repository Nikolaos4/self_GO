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
	Worker   int
	err      error
}

func countLines(filename string, id int) Result {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return Result{filename, 0, id, err}
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
	return Result{filename, count, id, nil}
}

func worker(id int, jobs <-chan string, results chan<- Result, wg *sync.WaitGroup) {
	for job := range jobs {
		results <- countLines(job, id)
	}
	wg.Done()
}

func stop(results chan Result, wg *sync.WaitGroup) {
	wg.Wait()
	close(results)
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan string, 6)
	{
		jobs <- "info.txt"
		jobs <- "error.txt"
		jobs <- "warn.txt"
		jobs <- "test1.txt"
		jobs <- "test2.txt"
		jobs <- "test3.txt"
	}
	close(jobs)

	results := make(chan Result, 6)
	//можно было написать wg.Add(3)
	wg.Add(1)
	go worker(1, jobs, results, &wg)
	wg.Add(1)
	go worker(2, jobs, results, &wg)
	wg.Add(1)
	go worker(3, jobs, results, &wg)
	go stop(results, &wg)

	for result := range results {
		fmt.Println(result)
	}
}
