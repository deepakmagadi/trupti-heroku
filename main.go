// You can edit this code!
// Click here and start typing.
package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func home(layout *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := layout.Execute(w, nil)
		if err != nil {
			log.Panic(err)
		}
	}
}

func main() {
	mux := http.NewServeMux()

	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	layout, err := template.ParseFiles("home.html.tmpl")
	if err != nil {
		log.Panic(err)
	}

	mux.HandleFunc("/", home(layout))

	log.Fatal(server.ListenAndServe())
}
