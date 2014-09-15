package main

import (
	"encoding/json"
	"github.com/hawx/riviera-admin/views"
	"flag"
	"log"
	"fmt"
	"net/http"
)

const HELP = `Usage: riviera-admin [options]

  An admin panel for riviera

    --port <num>     # Port to bind to (default: 8081)
    --riviera <url>  # Url to riviera (default: http://localhost:8080/)
    --help           # Display help message
`

var (
	port = flag.String("port", "8081", "")
	riviera = flag.String("riviera", "http://localhost:8080/", "")
	help = flag.Bool("help", false, "")
)

func main() {
	flag.Parse()

	if *help {
		fmt.Println(HELP)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(*riviera + "-/list")
		if err != nil {
			log.Print(err)
			w.WriteHeader(500)
			return
		}

		var list []string
		json.NewDecoder(resp.Body).Decode(&list)

		w.Header().Add("Content-Type", "text/html")
		views.Index.Execute(w, struct{
			Urls []string
			Riviera string
		}{list, *riviera})
	})

	log.Println("listening on port :" + *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
