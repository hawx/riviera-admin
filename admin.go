package main

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/hawx/persona"
	"github.com/hawx/riviera-admin/actions"
	"github.com/hawx/riviera-admin/views"
	"github.com/hawx/serve"

	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const HELP = `Usage: riviera-admin [options]

  An admin panel for riviera

    --port <num>        # Port to bind to (default: 8081)
    --socket <path>     # Serve using a unix socket instead
    --riviera <url>     # Url to riviera (default: http://localhost:8080/)

    --audience <host>   # Host and port site is running under (default: http://localhost:8081)
    --user <email>      # User who can access the admin panel
    --secret <str>      # String to use as cookie secret
    --path-prefix <p>   # Path prefix serving on

    --help              # Display help message
`

var (
	port       = flag.String("port", "8081", "")
	socket     = flag.String("socket", "", "")
	riviera    = flag.String("riviera", "http://localhost:8080/", "")
	audience   = flag.String("audience", "http://localhost:8081", "")
	user       = flag.String("user", "", "")
	secret     = flag.String("secret", "some-secret", "")
	pathPrefix = flag.String("path-prefix", "", "")
	help       = flag.Bool("help", false, "")
)

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	views.Login.Execute(w, struct{
		PathPrefix string
	}{*pathPrefix})
})

var List = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(*riviera + "-/list")
	if err != nil {
		log.Print("list", err)
		w.WriteHeader(500)
		return
	}

	var list []string
	json.NewDecoder(resp.Body).Decode(&list)

	w.Header().Add("Content-Type", "text/html")

	views.Index.Execute(w, struct {
		Url string
		PathPrefix string
		Feeds []string
	}{*audience, *pathPrefix, list})
})

var Subscribe = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	err := actions.Subscribe(*riviera, url)
	if err != nil {
		log.Println("subscribe:", err)
		w.WriteHeader(500)
		return
	}

	if r.FormValue("redirect") == "origin" {
		http.Redirect(w, r, url, 301)
		return
	}

	http.Redirect(w, r, "/", 301)
})

var Unsubscribe = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	err := actions.Unsubscribe(*riviera, r.FormValue("url"))
	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/", 301)
})

func main() {
	flag.Parse()

	if *help {
		fmt.Println(HELP)
		return
	}

	store := persona.NewStore(*secret)
	persona := persona.New(store, *audience, []string{*user})

	r := mux.NewRouter()

	r.Methods("GET").Path("/").Handler(persona.Switch(List, Login))
	r.Methods("GET").Path("/subscribe").Handler(persona.Protect(Subscribe))
	r.Methods("GET").Path("/unsubscribe").Handler(persona.Protect(Unsubscribe))
	r.Methods("POST").Path("/sign-in").Handler(persona.SignIn)
	r.Methods("GET").Path("/sign-out").Handler(persona.SignOut)

	http.Handle("/", r)

	serve.Serve(*port, *socket, context.ClearHandler(Log(http.DefaultServeMux)))
}
