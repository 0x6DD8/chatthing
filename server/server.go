package server

import (
	c "chatthing/common"
	"chatthing/controllers"
	"embed"
	"io/fs"
	"net/http"
)

func Start(port string, staticFiles embed.FS) error {
	setupRoutes(staticFiles)
	return http.ListenAndServe(port, nil)
}

func setupRoutes(staticFiles embed.FS) {
	setupStatic(staticFiles)

	http.Handle("/", c.SessionMiddleware(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		controllers.Root(w, req)
	})))

	http.Handle("/sse/", c.SessionMiddleware(http.HandlerFunc(controllers.SSE)))
}

func setupStatic(staticFiles embed.FS) {
	fs.WalkDir(staticFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			http.Handle("/"+path, c.SessionMiddleware(http.FileServer(http.FS(staticFiles))))
		}
		return nil
	})
}
