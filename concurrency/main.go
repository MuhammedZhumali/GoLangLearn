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
