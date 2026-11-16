package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func countLines(filename string) int {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return 0
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
	return count
}

func worker(jobs <-chan string, s *Stats, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		s.Add(job, countLines(job))
	}
}

//

type Stats struct {
	mu sync.Mutex
	total int
	ByFile map[string]int
}


func(s *Stats) Add(filename string, count int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.total += count
	s.ByFile[filename] = count
}


func main() {
	var wg sync.WaitGroup
	var stata Stats = Stats{ByFile: make(map[string]int)}
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

	//можно было написать wg.Add(3)
	wg.Add(1)
	go worker(jobs, &stata, &wg)
	wg.Add(1)
	go worker(jobs, &stata, &wg)
	wg.Add(1)
	go worker(jobs, &stata, &wg)
	
	wg.Wait()

	fmt.Println(stata.total)
	fmt.Println(stata.ByFile)
}

