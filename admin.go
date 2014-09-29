package main

import (
	"github.com/gorilla/context"
	"github.com/hawx/riviera-admin/actions"
	"github.com/hawx/riviera-admin/views"
	"github.com/hawx/wwwhat/persona"

	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

const HELP = `Usage: riviera-admin [options]

  An admin panel for riviera

    --port <num>       # Port to bind to (default: 8081)
    --socket <path>    # Serve using a unix socket instead
    --riviera <url>    # Url to riviera (default: http://localhost:8080/)

    --audience <host>  # Domain site is running under (default: localhost)
    --user <email>     # User who can access the admin panel
    --secret <str>     # String to use as cookie secret

    --help             # Display help message
`

var (
	port     = flag.String("port", "8081", "")
	socket   = flag.String("socket", "", "")
	riviera  = flag.String("riviera", "http://localhost:8080/", "")

	audience = flag.String("audience", "localhost", "")
	user     = flag.String("user", "", "")
	secret   = flag.String("secret", "some-secret", "")

	help     = flag.Bool("help", false, "")
)

var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	views.Login.Execute(w, struct{}{})
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
	views.Index.Execute(w, list)
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
	protect := persona.Protector(store, []string{*user})
	cond := persona.Conditional(store, []string{*user})

	http.Handle("/", cond(List, Login))
	http.Handle("/subscribe", protect(Subscribe))
	http.Handle("/unsubscribe", protect(Unsubscribe))
	http.Handle("/sign-in", persona.SignIn(store, *audience))
	http.Handle("/sign-out", persona.SignOut(store))

	if *socket == "" {
		log.Println("listening on port :" + *port)
		log.Fatal(http.ListenAndServe(":"+*port, context.ClearHandler(http.DefaultServeMux)))

	} else {
		l, err := net.Listen("unix", *socket)
		if err != nil {
			log.Fatal(err)
		}

		defer l.Close()

		go func() {
			log.Println("listening on", *socket)
			log.Fatal(http.Serve(l, context.ClearHandler(http.DefaultServeMux)))
		}()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)

		s := <-c
		log.Printf("caught %s: shutting down", s)
	}
}
