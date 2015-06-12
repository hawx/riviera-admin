package handlers

import (
	"log"
	"net/http"

	"hawx.me/code/riviera-admin/actions"
)

func Unsubscribe(opmlPath, pathPrefix string) http.Handler {
	return &unsubscribeHandler{opmlPath, pathPrefix}
}

type unsubscribeHandler struct {
	opmlPath   string
	pathPrefix string
}

func (h *unsubscribeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := actions.Unsubscribe(h.opmlPath, r.FormValue("url"))
	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, h.pathPrefix+"/", 301)
}
