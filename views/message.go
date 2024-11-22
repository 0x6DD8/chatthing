package components

import (
	"bytes"
	c "chatthing/common"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func MessageView(messages *c.MessageManager) Node {
	var messageCount = len(messages.Messages)
	var allMessages = make([]Node, messageCount)

	messages.Mu.Lock()
	defer messages.Mu.Unlock()

	for message := range messages.Messages {
		allMessages = append(allMessages,
			Div(
				B(Text(message.SessionID+": ")),
				Div(Text(message.Message), Class("message-content")),
				Class("message"),
			),
		)
	}
	allMessages = append(allMessages, Class("message-container"))

	return Div(allMessages...)
}

func MessageViewAsString(messages *c.MessageManager) string {
	var buf bytes.Buffer
	err := MessageView(messages).Render(&buf)
	if err != nil {
		return ""
	}
	return buf.String()
}
