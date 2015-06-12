package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"hawx.me/code/mux"
	"hawx.me/code/persona"
	"hawx.me/code/riviera-admin/handlers"
	"hawx.me/code/serve"
)

const HELP = `Usage: riviera-admin [options] FILE

  An admin panel for riviera.

    --port <num>        # Port to bind to (default: 8081)
    --socket <path>     # Serve using a unix socket instead

    --audience <host>   # Host and port site is running under (default: http://localhost:8081)
    --user <email>      # User who can access the admin panel
    --secret <str>      # String to use as cookie secret
    --path-prefix <p>   # Path prefix serving on

    --help              # Display help message
`

var (
	port       = flag.String("port", "8081", "")
	socket     = flag.String("socket", "", "")
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

func main() {
	flag.Parse()

	if *help || flag.NArg() == 0 {
		fmt.Println(HELP)
		return
	}

	opmlPath := flag.Arg(0)

	store := persona.NewStore(*secret)
	persona := persona.New(store, *audience, []string{*user})

	http.Handle("/", mux.Method{
		"GET": persona.Switch(
			handlers.List(opmlPath, *audience, *pathPrefix),
			handlers.Login(*pathPrefix),
		),
	})
	http.Handle("/subscribe", persona.Protect(mux.Method{
		"GET": handlers.Subscribe(opmlPath, *pathPrefix),
	}))
	http.Handle("/unsubscribe", persona.Protect(mux.Method{
		"GET": handlers.Unsubscribe(opmlPath, *pathPrefix),
	}))
	http.Handle("/sign-in", mux.Method{
		"POST": persona.SignIn,
	})
	http.Handle("/sign-out", mux.Method{
		"GET": persona.SignOut,
	})

	serve.Serve(*port, *socket, Log(http.DefaultServeMux))
}
