package main

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

type State struct {
	//mu    sync.Mutex
	count int32
}

func (s *State) setState(i int) {
	atomic.AddInt32(&s.count, int32(i))
	//s.mu.Lock()
	//defer s.mu.Unlock()
	//s.count = i
}

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "username", "Emirhangg")
	userId, err := fetchuserId(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Response took %v -> %+v\n", time.Since(start), userId)

}

func fetchuserId(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	type Result struct {
		userId string
		err    error
	}

	val := ctx.Value("username")
	fmt.Println("username:", val)

	resultch := make(chan Result, 1)

	go func() {
		response, err := thirdPartyHttpCall()
		resultch <- Result{
			userId: response,
			err:    err,
		}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-resultch:
		return result.userId, result.err
	}
}

func thirdPartyHttpCall() (string, error) {
	time.Sleep(time.Millisecond * 100)
	return "user id 1", nil
}

// Keeping Server alive

// type Server struct {
// 	quitch chan struct{}
// 	msgch  chan string
// }

// func newServer() *Server {
// 	return &Server{
// 		quitch: make(chan struct{}),
// 		msgch:  make(chan string, 128),
// 	}
// }

// func (s *Server) sendMessage(msg string) {
// 	s.msgch <- msg
// }

// func (s *Server) start() {
// 	fmt.Println("Server is starting...")
// 	s.loop()
// }

// func (s *Server) stop() {
// 	close(s.quitch)
// }

// func (s *Server) loop() {
// mainloop:
// 	for {
// 		select {
// 		case <-s.quitch:
// 			fmt.Println("quit")
// 			break mainloop
// 		case msg := <-s.msgch:
// 			s.handleMessage(msg)
// 		}

// 	}
// 	fmt.Println("quitting gracefully")
// }

// func (s *Server) handleMessage(msg string) {
// 	fmt.Println("Received message:", msg)
// }
