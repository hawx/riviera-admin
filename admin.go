package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"

	"hawx.me/code/mux"
	"hawx.me/code/riviera-admin/handlers"
	"hawx.me/code/serve"
	"hawx.me/code/uberich"
)

const HELP = `Usage: riviera-admin [options] FILE

  An admin panel for riviera.

    --port <num>        # Port to bind to (default: 8081)
    --socket <path>     # Serve using a unix socket instead
    --settings <path>   # Path to settings (default: ./settings.toml)

    --help              # Display help message
`

type Conf struct {
	Secret     string
	URL        string
	PathPrefix string

	Uberich struct {
		AppName    string
		AppURL     string
		UberichURL string
		Secret     string
	}
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	var (
		settingsPath = flag.String("settings", "./settings.toml", "")
		port         = flag.String("port", "8081", "")
		socket       = flag.String("socket", "", "")
	)
	flag.Usage = func() { fmt.Println(HELP) }
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println(HELP)
		return
	}

	opmlPath := flag.Arg(0)

	var conf *Conf
	if _, err := toml.DecodeFile(*settingsPath, &conf); err != nil {
		log.Fatal("toml: ", err)
	}

	store := uberich.NewStore(conf.Secret)
	uberich := uberich.NewClient(conf.Uberich.AppName, conf.Uberich.AppURL, conf.Uberich.UberichURL, conf.Uberich.Secret, store)

	shield := func(h http.Handler) http.Handler {
		return uberich.Protect(h, http.NotFoundHandler())
	}

	http.Handle("/", mux.Method{
		"GET": uberich.Protect(
			handlers.List(opmlPath, conf.URL, conf.PathPrefix),
			handlers.Login(conf.URL, conf.PathPrefix),
		),
	})
	http.Handle("/subscribe", shield(mux.Method{
		"GET": handlers.Subscribe(opmlPath, conf.PathPrefix),
	}))
	http.Handle("/unsubscribe", shield(mux.Method{
		"GET": handlers.Unsubscribe(opmlPath, conf.PathPrefix),
	}))
	http.Handle("/sign-in", uberich.SignIn("/"))
	http.Handle("/sign-out", uberich.SignOut("/"))

	serve.Serve(*port, *socket, Log(http.DefaultServeMux))
}
