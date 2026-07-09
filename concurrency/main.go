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

	statusCh := make(chan string, 1)
	statusCh <- "ready"
	close(statusCh)
	value, ok := <-statusCh
	fmt.Println(value, ok)
	value, ok = <-statusCh
	fmt.Println(value, ok)

	firstCh := make(chan string)
	secondCh := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		firstCh <- "first"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		secondCh <- "second"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-firstCh:
			fmt.Println(msg1)
		case msg2 := <-secondCh:
			fmt.Println(msg2)
		}
	}

	resultCh := make(chan string, 1)
	go func() {
		time.Sleep(200 * time.Millisecond)
		resultCh <- "result"
	}()
	select {
	case res := <-resultCh:
		fmt.Println(res)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("timeout")
	}

	counter := 0
	var mu sync.Mutex
	var wg2 sync.WaitGroup
	wg2.Add(2)
	go func() {
		defer wg2.Done()
		for i := 0; i < 1000; i++ {
			func(){
			mu.Lock()
			defer mu.Unlock() // Ensure the mutex is unlocked even if panic occurs
			counter++
			}()
		}
	}()
	go func() {
		defer wg2.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock() // default behavior, but can lead to deadlocks if not handled properly
			counter++
			mu.Unlock()
		}
	}()

	wg2.Wait()
	fmt.Println("Counter:", counter)
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
