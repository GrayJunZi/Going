package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} // 0 bytes
	msgch  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("server starting")
	s.loop()
}

func (s *Server) quit() {
	s.quitch <- struct{}{}
}

func (s *Server) loop() {
mainloop:
	for {
		select {
		case <-s.quitch:
			break mainloop
		case msg := <-s.msgch:
			s.handleMessage(msg)
		}
	}
	fmt.Println("server is shutting down gracefully")
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("we received a message:", msg)
}

func main() {
	server := newServer()

	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()

	server.start()
}
