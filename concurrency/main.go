package main

import (
	"fmt"

	"sync"

	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		printNumbers()
	}()
	go func() {
		defer wg.Done()
		printLetters()
	}()
	wg.Wait()
	fmt.Println("main finished")

	// Channels
	ch := make(chan string)
	go func() {
		ch <- "task completed"
	}()
	message := <-ch
	fmt.Println(message)

	ch2 := make(chan int)
	go calculateSquare(7, ch2)
	result := <-ch2
	fmt.Println(result)

	bufferedCh := make(chan int, 2)
	bufferedCh <- 10
	bufferedCh <- 20
	//bufferedCh <- 30 // error
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)
	bufferedCh <- 30 // fine
	fmt.Println(<-bufferedCh)

	numbers := make(chan int, 3)
	numbers <- 10
	numbers <- 20
	numbers <- 30
	close(numbers)
	for number := range numbers {
		fmt.Println(number)
	}
	// numbers <- 40 // Panic

}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(i)
	}
}

func printLetters() {
	str := []string{"A", "B", "C", "D", "E"}
	for i := 0; i < len(str); i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(str[i])
	}
}

func calculateSquare(number int, ch chan<- int) {
	number *= number
	ch <- number
}
