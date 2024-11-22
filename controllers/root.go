package controllers

import (
	v "chatthing/views"
	"net/http"
)

func Root(w http.ResponseWriter, req *http.Request) {
	var err = v.Root().Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
