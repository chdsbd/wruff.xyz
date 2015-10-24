package main

import (
	"net/http"
	"text/template"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gopkg.in/tylerb/graceful.v1"
)

type Message struct {
	Title string
	data  string
}

var templates = template.Must(template.ParseFiles("index.html", "404.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "index.html", Message{Title: "index"})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "404.html", Message{})
}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	//serve static files inside the public folder ( make sure to prefix)
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	r.HandleFunc("/", IndexHandler)
	n := negroni.Classic()
	n.UseHandler(r)
	graceful.Run(":8000", 10*time.Second, n)
}

func renderPage(w http.ResponseWriter, templateName string, data interface{}) {
	err := templates.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
