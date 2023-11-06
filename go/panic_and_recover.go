package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{1}
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("catch panic")
		}
	}()
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("catch panic")
			}
		}()
		fmt.Println(arr[1])
	}()
	time.Sleep(time.Second)
	fmt.Println(arr[0])
}
