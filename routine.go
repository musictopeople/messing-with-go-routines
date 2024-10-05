package main

import (
	"fmt"
	"time"
)

func routine(line string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(line, " : ", i)
	}
}

func runRoutines() {
	fmt.Println("\nRunning Go Routine Example: \n")
	go routine("concurrent")
	routine("sync      ")
	fmt.Println("")
}
