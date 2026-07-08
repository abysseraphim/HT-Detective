package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	jobs := make(chan string, 100)

	var wg sync.WaitGroup

	workers := flag.Int("w", 10, "number of workers")
	flag.Parse()

	fmt.Fprintf(os.Stderr, "Started Job With %d Threads(workers)... \n", *workers)

	wg.Add(*workers)

	for i := 0; i < *workers; i++ {
		go worker(jobs, client, &wg)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		jobs <- line
	}
	err := scanner.Err()
	if err != nil {
		fmt.Println("error occured:", err)
	}

	close(jobs)

	wg.Wait()
}
