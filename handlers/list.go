package handlers

import (
	"log"
	"net/http"
	"sort"
	"strings"

	"hawx.me/code/riviera-admin/views"
	"hawx.me/code/riviera/subscriptions/opml"
)

func List(opmlPath, url string) http.Handler {
	return &listHandler{opmlPath, url}
}

type listHandler struct {
	opmlPath string
	url      string
}

func (h *listHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	outline, err := opml.Load(h.opmlPath)
	if err != nil {
		log.Println("list", err)
		w.WriteHeader(500)
		return
	}

	var list feeds
	for _, line := range outline.Body.Outline {
		if line.Type == "rss" {
			f := feed{
				FeedUrl:         line.XMLURL,
				WebsiteUrl:      line.HTMLURL,
				FeedTitle:       line.Title,
				FeedDescription: line.Description,
			}

			if f.FeedTitle == "" {
				f.FeedTitle = f.FeedUrl
			}

			list = append(list, f)
		}
	}

	sort.Sort(list)

	w.Header().Add("Content-Type", "text/html")

	views.Index.Execute(w, struct {
		Url   string
		Feeds []feed
	}{h.url, list})
}

type feed struct {
	FeedUrl         string
	WebsiteUrl      string
	FeedTitle       string
	FeedDescription string
}

type feeds []feed

func (fs feeds) Len() int      { return len(fs) }
func (fs feeds) Swap(i, j int) { fs[i], fs[j] = fs[j], fs[i] }
func (fs feeds) Less(i, j int) bool {
	return strings.ToLower(fs[i].FeedTitle) < strings.ToLower(fs[j].FeedTitle)
}
