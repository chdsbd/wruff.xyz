package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/sbdchd/wruff.xyz/apis"
	"gopkg.in/tylerb/graceful.v1"
)

type Message struct {
	Title string
	data  string
}

var message string = "Wruff Wruff!"

var templates = template.Must(template.ParseFiles("index.html", "404.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == "POST" {
		if email := r.FormValue("InputEmail"); email != "" {
			err = apis.SendEmail(email, message)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		if phoneNumber := r.FormValue("InputPhoneNumber"); phoneNumber != "" {
			err = apis.SendSMS(phoneNumber, message)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			err = apis.SendCall(phoneNumber)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		if twitterHandle := r.FormValue("InputTwitterHandle"); twitterHandle != "" {
			err = apis.SendTweet(twitterHandle, message)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		if yoUsername := r.FormValue("InputYoUsername"); yoUsername != "" {
			err = apis.SendYo(yoUsername)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

	}
	renderPage(w, "index.html", Message{})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "404.html", Message{})
}

func main() {
	err := checkEnv()
	if err != nil {
		log.Fatal(err)
	}
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

func checkEnv() error {
	env := []string{"mg_apiKey",
		"mg_publicApiKey",
		"yo_apikey",
		"Twilio_AccoutSid",
		"Twilio_AuthToken",
		"Twilio_Number",
		"Twilio_TestPhone",
		"reddit_user",
		"reddit_password",
		"reddit_clientid",
		"reddit_secret",
		"twitter_ConsumerKey",
		"twitter_ConsumerSecret",
		"twitter_AccessToken",
		"twitter_AcessTokenSecret",
		"twitter_TestUsername",
		"yo_username",
	}
	for _, v := range env {
		if i := os.Getenv(v); i == "" {
			return errors.New(fmt.Sprintf("env %s not set", v))
			break
		}
	}
	return nil

}
