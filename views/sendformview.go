package components

import (
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func SendFormView() Node {
	return Div(
		ID("send-message-form"),
		Form(
			Class("send-message-form"),
			Input(Class("message-input"), Name("message"), Attr("required")),
			Button(Text("send"), hx.Post("/sse/send"), hx.Target("#send-message-form"), hx.Swap("outerHTML"), Class("send-msg")),
		),
	)
}
