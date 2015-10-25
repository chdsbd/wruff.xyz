# wruff.xyz [![Build Status](https://travis-ci.org/chdsbd/wruff.xyz.svg?branch=master)](https://travis-ci.org/chdsbd/wruff.xyz)
Send all the messages! :smiling_imp:

## What is this?

This is a project for [hackumass](http://hackumass.com).

*Note:* The website is not enabled for obvious reasons.

## Using

Webserver
- [gorrilla mux](https://github.com/gorilla/mux)
- [negroni](https://github.com/codegangsta/negroni)
- [graceful.v1](https://github.com/tylerb/graceful)

Frontend
- [Bootstrap](http://getbootstrap.com/)
- [Font Awesome](http://fontawesome.io/)

APIs
- [Mailgun](https://documentation.mailgun.com/api_reference.html)
- [Twilio](https://www.twilio.com/docs)
- [Yo](https://dev.justyo.co)
- [Twitter](https://dev.twitter.com/rest/public)


## Installation & Setup

```bash
git clone https://github.com/chdsbd/wruff.xyz
```
```bash
go run main.go
```

Setup all the Enviornmental variables

- ```Twilio_AccoutSid```
- ```Twilio_AuthToken```
- ```Twilio_Number```
- ```Twilio_TestPhone```
- ```mg_apiKey```
- ```mg_publicAPIKey```
- ```yo_apikey```
- ```yo_username```
- ```twitter_ConsumerKey```
- ```twitter_ConsumerSecret```
- ```twitter_AccessToken```
- ```twitter_AcessTokenSecret```
- ```twitter_TestUsername```
