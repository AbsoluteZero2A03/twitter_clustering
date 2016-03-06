package main

import (
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"log"
	"net/http"
	"os"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(twitter.AuthorizeEndpoint))
	config := oauth1.Config{
		ConsumerKey:    os.Getenv("TWITTER_APP_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("TWITTER_APP_CONSUMER_SECRET"),
		CallbackURL:    "http://127.0.0.1:9002/callback",
		Endpoint:       twitter.AuthorizeEndpoint,
	}
	requestToken, _, err := config.RequestToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	authorizationURL, err := config.AuthorizationURL(requestToken)
	fmt.Println(authorizationURL)

	auth_redirect := GenerateRedirect(authorizationURL.String())

	http.HandleFunc("/", auth_redirect)
	http.HandleFunc("/callback", Callback)
	err = http.ListenAndServe("127.0.0.1:9002", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func GenerateRedirect(urlStr string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, urlStr, http.StatusFound)
	}
}

func Callback(w http.ResponseWriter, r *http.Request) {
	requestToken, verifier, err := oauth1.ParseAuthorizationCallback(r)
	if err != nil {
		log.Fatal("Callback: ", err)
	}

	fmt.Println(requestToken)
	fmt.Println(verifier)

	http.Redirect(w, r, "http://www.google.com", http.StatusFound)
}
