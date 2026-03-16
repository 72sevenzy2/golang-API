package service

import (
	"fmt"
	"sync"
)

type Greeter interface {
	Greet(name string) (string, error)
}

type GreetCounter struct {
	count int
	mu    sync.Mutex
}

func (s *GreetCounter) Greet(name string) (string, int, error) {
	var count int
	if name == "" {
		return "", 0, fmt.Errorf("name cannot be empty")
	}

	s.mu.Lock()
	s.count = s.count + 1
	count = s.count
	s.mu.Unlock()

	return fmt.Sprintf("welcome back %s, greet number %d", name, count), count, nil
}