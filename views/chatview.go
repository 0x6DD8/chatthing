package components

import (
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func ChatView() Node {
	return Div(
		Div(Attr("sse-swap", "join"), Class("user-count")),
		Div(Attr("sse-swap", "message"), Class("message-list")),
		hx.Ext("sse"),
		Attr("sse-connect", "/sse/connect"),
	)
}
