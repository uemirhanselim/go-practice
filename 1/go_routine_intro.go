package main

import (
	"fmt"
	"time"
)

func delayedHello() {
	fmt.Println("Starting delayedHello")
	time.Sleep(2 * time.Second)
	fmt.Println("Hello, World!")
}

func main() {
	msgch := make(chan string, 128)

	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	close(msgch)

	for {
		msg, ok := <-msgch

		if !ok {
			break
		}

		fmt.Println(msg)
	}

	// for msg := range msgch {
	// 	fmt.Println(msg)
	// }
}
