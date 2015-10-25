package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/chdsbd/wruff.xyz/apis"
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
	message := fmt.Sprintf("Wruff Wruff! %s", time.Now())
	var err error
	if r.Method == "POST" {
		if email := r.FormValue("InputEmail"); email != "" {
			go func() {
				err = apis.SendEmail(email, message)
				if err != nil {
					log.Println("Email", err)
				}
			}()
		} else {
			log.Println("Email Missing")
		}
		if phoneNumber := r.FormValue("InputPhoneNumber"); phoneNumber != "" {
			go func() {
				err = apis.SendSMS(phoneNumber, message)
				if err != nil {
					log.Println("SMS", err)
				}
			}()
			go func() {
				err = apis.SendCall(phoneNumber)
				if err != nil {
					log.Println("Voice", err)
				}
			}()

		} else {
			log.Println("Phone Number Missing")
		}
		if twitterHandle := r.FormValue("InputTwitterHandle"); twitterHandle != "" {
			go func() {
				err = apis.SendTweet(twitterHandle, message)
				if err != nil {
					log.Println(err)
				}
			}()
		}
		if yoUsername := r.FormValue("InputYoUsername"); yoUsername != "" {
			go func() {
				err = apis.SendYo(yoUsername)
				if err != nil {
					log.Println("Yo", err)
				}
			}()
		} else {
			log.Println("Yo Username Missing")
		}
		log.Println("SUCCESS!")
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
