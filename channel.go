package main

import "fmt"

func characters(s []string, c chan string) {
	concat := "a"
	for _, v := range s {
		concat += v
	}
	c <- concat
}

func runChannels() {
	fmt.Println("\nRunning Go Channels Example: \n")
	s := []string{"b", "c"}
	c := make(chan string)
	bufferedChannel := make(chan string, 3)
	go characters(s, c)
	go characters(s, bufferedChannel)
	x, y := <-c, <-bufferedChannel
	fmt.Println(x, y, x+y)
	fmt.Println("")
}
