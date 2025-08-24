package main

import (
	"fmt"
	"time"
)

func sendMessage1(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Message from Goroutine 1"
}

func sendMessage2(ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- "Message from Goroutine 2"
}

func startSelect() {
	// Create two channels for communication
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start two goroutines that will send messages
	go sendMessage1(ch1)
	go sendMessage2(ch2)

	// Use select to listen on multiple channels
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	case <-time.After(3 * time.Second): // Timeout case to avoid waiting indefinitely
		fmt.Println("Timeout: No messages received")
	}
}
