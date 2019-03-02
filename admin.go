package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"

	"hawx.me/code/indieauth"
	"hawx.me/code/indieauth/sessions"
	"hawx.me/code/mux"
	"hawx.me/code/riviera-admin/handlers"
	"hawx.me/code/serve"
)

const help = `Usage: riviera-admin [options] FILE

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
	Me         string
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
	flag.Usage = func() { fmt.Println(help) }
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println(help)
		return
	}

	opmlPath := flag.Arg(0)

	var conf *Conf
	if _, err := toml.DecodeFile(*settingsPath, &conf); err != nil {
		log.Fatal("toml: ", err)
	}

	auth, err := indieauth.Authentication(conf.URL, conf.URL+"/callback")
	if err != nil {
		log.Fatal(err)
	}

	session, err := sessions.New(conf.Me, conf.Secret, auth)
	if err != nil {
		log.Fatal(err)
	}
	session.Root = conf.PathPrefix

	http.Handle("/", mux.Method{
		"GET": session.Choose(
			handlers.List(opmlPath, conf.URL, conf.PathPrefix),
			handlers.Login(conf.URL, conf.PathPrefix),
		),
	})
	http.Handle("/subscribe", session.Shield(mux.Method{
		"GET": handlers.Subscribe(opmlPath, conf.PathPrefix),
	}))
	http.Handle("/unsubscribe", session.Shield(mux.Method{
		"GET": handlers.Unsubscribe(opmlPath, conf.PathPrefix),
	}))
	http.Handle("/sign-in", session.SignIn())
	http.Handle("/callback", session.Callback())
	http.Handle("/sign-out", session.SignOut())

	serve.Serve(*port, *socket, Log(http.DefaultServeMux))
}
