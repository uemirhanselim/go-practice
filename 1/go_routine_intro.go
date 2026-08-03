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
	done := make(chan bool)

	printerch := make(chan string)

	go func() {
		value := <-printerch
		fmt.Println(value)
		done <- true
	}()

	printerch <- "foo"
	<-done
	//go delayedHello()
}
