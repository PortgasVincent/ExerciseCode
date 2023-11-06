package _go

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("done")
				return
			default:
				fmt.Println("running")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	stop <- true
	time.Sleep(5 * time.Second)
	return
}

// for range 用法   https://tour.go-zh.org/concurrency/4
