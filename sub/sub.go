package sub

import (
	"sync"
)

type Sub struct {
	Topics map[string][]chan string
	Mu     sync.Mutex
}

func NewSub() *Sub {
	return &Sub{
		Topics: make(map[string][]chan string),
	}
}

func (s *Sub) Subscribe(topic string) chan string {
	ch := make(chan string)
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Topics[topic] = append(s.Topics[topic], ch)
	return ch
}

func (s *Sub) Unsubscribe(topic string, ch chan string) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	subscribers := s.Topics[topic]
	for i, sub := range subscribers {
		if sub == ch {
			// Remove the subscriber
			s.Topics[topic] = append(subscribers[:i], subscribers[i+1:]...)
			close(sub)
			break
		}
	}
}
