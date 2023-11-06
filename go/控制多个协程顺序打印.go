package main

import (
	"fmt"
	"time"
)

func main() {
	var a, b = make(chan int), make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("A")
			a <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(a)
	}()
	go func() {
		for range a {
			fmt.Println("B")
		}
		close(b)
	}()
	for range b {
		fmt.Println("C")
	}
}
