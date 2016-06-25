package handlers

import (
	"net/http"

	"hawx.me/code/riviera-admin/views"
)

func Login(url, pathPrefix string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		views.Login.Execute(w, struct {
			Url        string
			PathPrefix string
		}{url, pathPrefix})
	})
}
