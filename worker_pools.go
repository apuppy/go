package main

import "time"
import "fmt"

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("workd id", id, "processing job", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		<- results
	}
}