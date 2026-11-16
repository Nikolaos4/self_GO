package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"math/rand"
)

type Result struct {
	Filename string
	count    int
	Worker   int
	err      error
}

func countLines(filename string, id int) Result {
	time.Sleep(time.Duration(rand.Intn(3000))*time.Millisecond)
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

func worker(id int, jobs <-chan string, results chan<- Result, stop_ch <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
			case job, ok := <- jobs:
				if ok == false {
					return
				}
				results <- countLines(job, id)
			case <-stop_ch:
				fmt.Printf("Worker %d stopped \n", id)
				return
		}
	}
}

func stop(results chan Result, wg *sync.WaitGroup) {
	wg.Wait()
	close(results)
}

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})

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
	go worker(1, jobs, results, done, &wg)
	wg.Add(1)
	go worker(2, jobs, results, done, &wg)
	wg.Add(1)
	go worker(3, jobs, results, done, &wg)
	go stop(results, &wg)

	timeout := time.After(2 * time.Second)
	Cycle:
	for {
		select {
			case current_res, ok := <-results:
				if ok == false {
					break Cycle
				}
				fmt.Println(current_res)
			case time_flag := <-timeout:
				fmt.Println("Timeout: ", time_flag)
				close(done)
				break Cycle
		}
	}
	time.Sleep(4*time.Second)
}

