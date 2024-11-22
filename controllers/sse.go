package controllers

import (
	c "chatthing/common"
	h "chatthing/views"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func SSE(w http.ResponseWriter, req *http.Request) {
	c.ManageSession(w, req)

	switch req.URL.Path {
	case "/sse/connect":
		handleConnect(w, req)
	case "/sse/send":
		handleSend(w, req)
	default:
		http.NotFound(w, req)
	}
}

var clientManager = c.ClientManager{
	Clients: make(map[*c.Client]bool),
}

var messageManager = c.MessageManager{
	Messages: make(map[*c.Message]bool),
}

func handleConnect(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	var client = &c.Client{W: w, Req: req}
	clientManager.AddClient(client)

	var flusher, ok = w.(http.Flusher)
	if !ok {
		http.Error(w, "err", http.StatusInternalServerError)
		return
	}

	SendEvent("join", strconv.Itoa(len(clientManager.Clients)), client)

	var renderedMessages = h.MessageViewAsString(&messageManager)
	SendEvent("message", renderedMessages, client)

	var notify = req.Context().Done()
	for {
		select {
		case <-notify:
			clientManager.RemoveClient(client)
			return
		default:
			flusher.Flush()
			SendComment("heartbeat", client)
			time.Sleep(time.Second)
		}
	}
}

func handleSend(w http.ResponseWriter, req *http.Request) {
	var message = req.FormValue("message")
	var sessionID = c.GetSessionID(req)

	h.SendFormView().Render(w)

	if message == "" {
		return
	}

	messageManager.AddMessage(&c.Message{SessionID: sessionID, Message: message})

	var renderedMessages = h.MessageViewAsString(&messageManager)
	BroadcastMessage(renderedMessages)

}

func BroadcastMessage(message string) {
	clientManager.Mu.Lock()
	defer clientManager.Mu.Unlock()
	for client := range clientManager.Clients {
		SendEvent("message", message, client)
	}
}

func SendEvent(event string, data string, client *c.Client) {
	var eventToWrite = "event: " + event + "\n"
	var dataToWrite = "data: " + data + "\n\n"
	sendAndFlush(client.W, eventToWrite+dataToWrite)
}

func SendComment(comment string, client *c.Client) {
	var commentToWrite = ";" + comment + "\n\n"
	sendAndFlush(client.W, commentToWrite)
}

func sendAndFlush(w http.ResponseWriter, comment string) {
	var _, err = fmt.Fprintf(w, comment)
	if err != nil {
		var rc = http.NewResponseController(w)
		err = rc.Flush()
	}
}
