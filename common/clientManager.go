package common

import (
	"net/http"
	"sync"
)

func (m *ClientManager) AddClient(client *Client) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	m.Clients[client] = true
}

func (m *ClientManager) RemoveClient(client *Client) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	delete(m.Clients, client)
}

type Client struct {
	W         http.ResponseWriter
	Req       *http.Request
	SessionID string
}

type ClientManager struct {
	Clients map[*Client]bool
	Mu      sync.Mutex
}
