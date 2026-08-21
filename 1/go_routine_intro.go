package main

import (
	"fmt"
)

type Server struct {
	quitch chan struct{}
	msgch  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string),
	}
}

func (s *Server) start() {
	fmt.Println("Server is starting...")
}

func (s *Server) loop() {
	for {
		select {
		case <-s.quitch:
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default:
		}

	}
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("Received message:", msg)
}

func main() {

}
