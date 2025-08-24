package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func routinetimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jobs := make(chan string)

	go func() {
		fmt.Println("Sleeping... 6 seconds")
		time.Sleep(6 * time.Second)
		jobs <- "Long routine"
	}()

	go func() {
		fmt.Println("Sleeping... 4 seconds")
		time.Sleep(4 * time.Second)
		jobs <- "Short routine"
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context timeout before long routine")
			os.Exit(0)
		case res := <-jobs:
			// handle val
			fmt.Println("Result: ", res)
		}
	}
}
