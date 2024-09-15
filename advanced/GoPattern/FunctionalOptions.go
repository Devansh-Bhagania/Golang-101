// Provides a way to set optional parameters when creating an object.

package main

import (
	"fmt"
	"time"
)

type Server struct {
	Addr    string
	Port    int
	Timeout time.Duration
}

type Option func(*Server)

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func NewServer(addr string, opts ...Option) *Server {
	s := &Server{
		Addr:    addr,
		Port:    8080,             // default port
		Timeout: 30 * time.Second, // default timeout
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func main() {
	server := NewServer("localhost", WithPort(9090), WithTimeout(60*time.Second))
	fmt.Printf("%+v\n", server)
}
