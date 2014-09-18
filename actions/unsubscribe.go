package actions

import (
	"net/http"
	"net/url"
)

func Unsubscribe(riviera, page string) error {
	_, err := http.PostForm(riviera+"-/unsubscribe", url.Values{"url": {page}})
	return err
}
