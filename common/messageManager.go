package common

import (
	"sync"
)

func (m *MessageManager) AddMessage(message *Message) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	m.Messages[message] = true
}

func (m *MessageManager) RemoveMessage(message *Message) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	delete(m.Messages, message)
}

type Message struct {
	Message   string
	SessionID string
}

type MessageManager struct {
	Messages map[*Message]bool
	Mu       sync.Mutex
}
