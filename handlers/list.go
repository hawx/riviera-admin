package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"hawx.me/code/riviera-admin/views"
)

func List(riviera, audience, pathPrefix string) http.Handler {
	return &listHandler{riviera, audience, pathPrefix}
}

type listHandler struct {
	riviera    string
	audience   string
	pathPrefix string
}

func (h *listHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", h.riviera+"river/meta", nil)
	if err != nil {
		log.Println("list", err)
		w.WriteHeader(500)
		return
	}

	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("list", err)
		w.WriteHeader(500)
		return
	}

	var list []feed
	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		log.Println("list", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "text/html")

	views.Index.Execute(w, struct {
		Url        string
		PathPrefix string
		Feeds      []feed
	}{h.audience, h.pathPrefix, list})
}

type feed struct {
	FeedUrl         string
	WebsiteUrl      string
	FeedTitle       string
	FeedDescription string
	Status          string
}
