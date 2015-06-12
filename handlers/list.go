package handlers

import (
	"log"
	"net/http"

	"hawx.me/code/riviera-admin/views"
	"hawx.me/code/riviera/subscriptions/opml"
)

func List(opmlPath, audience, pathPrefix string) http.Handler {
	return &listHandler{opmlPath, audience, pathPrefix}
}

type listHandler struct {
	opmlPath   string
	audience   string
	pathPrefix string
}

func (h *listHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	outline, err := opml.Load(h.opmlPath)
	if err != nil {
		log.Println("list", err)
		w.WriteHeader(500)
		return
	}

	var list []feed
	for _, line := range outline.Body.Outline {
		if line.Type == "rss" {
			f := feed{
				FeedUrl:         line.XmlUrl,
				WebsiteUrl:      line.HtmlUrl,
				FeedTitle:       line.Title,
				FeedDescription: line.Description,
			}

			if f.FeedTitle == "" {
				f.FeedTitle = f.FeedUrl
			}

			list = append(list, f)
		}
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
}
