package handlers

import (
	"log"
	"net/http"

	"hawx.me/code/riviera-admin/actions"
)

func Subscribe(opmlPath, pathPrefix string) http.Handler {
	return &subscribeHandler{opmlPath, pathPrefix}
}

type subscribeHandler struct {
	opmlPath   string
	pathPrefix string
}

func (h *subscribeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	err := actions.Subscribe(h.opmlPath, url)
	if err != nil {
		log.Println("subscribe:", err)
		w.WriteHeader(500)
		return
	}

	if r.FormValue("redirect") == "origin" {
		http.Redirect(w, r, url, 301)
		return
	}

	http.Redirect(w, r, h.pathPrefix+"/", 301)
}
