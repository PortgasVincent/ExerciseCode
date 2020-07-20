package main

import (
	"fmt"
)

func maintest() {
	m := map[string]string{
		"aa":"aa",
	}
	if m["aa"] {
		fmt.Println("bool")
	}
	if m["aa"] == "aa" {
		fmt.Println("string")
	}
}