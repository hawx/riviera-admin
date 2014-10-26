package actions

import (
	"github.com/PuerkitoBio/goquery"

	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func Subscribe(riviera, page string) error {
	feed, err := getFeed(page)
	if err != nil {
		return err
	}

	_, err = http.PostForm(riviera+"-/subscribe", url.Values{"url": []string{feed}})
	return err
}

func getFeed(page string) (string, error) {
	resp, err := http.Get(page)
	if err != nil {
		return "", err
	}

	contentType := resp.Header.Get("Content-Type")
	if isFeedType(contentType) {
		return page, nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	for _, node := range doc.Find("link").Nodes {
		var rel, href, typ string
		for _, a := range node.Attr {
			switch a.Key {
			case "rel":
				rel = a.Val
			case "href":
				href = a.Val
			case "type":
				typ = a.Val
			}
		}

		if rel == "alternate" && isFeedType(typ) {
			if strings.HasPrefix(href, "/") {
				u, err := url.Parse(page)
				if err != nil {
					return "", err
				}
				u.Path = href
				return u.String(), nil
			} else {
				return href, nil
			}
		}
	}

	return "", errors.New("no feed found on page")
}

var contentTypes = []*regexp.Regexp{
	regexp.MustCompile("^application/rss\\+xml"),
	regexp.MustCompile("^application/atom\\+xml"),
	regexp.MustCompile("^application/xml"),
	regexp.MustCompile("^text/xml"),
}

func isFeedType(contentType string) bool {
	for _, pat := range contentTypes {
		if pat.Match([]byte(contentType)) {
			return true
		}
	}

	return false
}
