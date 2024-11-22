package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Root() Node {
	return HTML(
		Head(
			Meta(Charset("utf-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
			Title("Chatthing"),
			Script(Src("static/htmx.js")),
			Script(Src("static/htmx-ext-sse.js")),
			Link(Rel("stylesheet"), Href("static/styles.css")),
		),
		Body(
			Header(
				A(
					H1(Text("Chatthing")), Href("/"),
				),
			),
			Main(
				ChatView(),
				SendFormView(),
			),
		),
	)
}
