package handlers

import (
	"net/http"

	"hawx.me/code/riviera-admin/views"
)

func Login(pathPrefix string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		views.Login.Execute(w, struct {
			PathPrefix string
		}{pathPrefix})
	})
}
