package main

import "sync/atomic"

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
