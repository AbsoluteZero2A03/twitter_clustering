package main

import (
	"fmt"
    "io/ioutil"
    "encoding/json"
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
	requestToken, requestSecret, err := config.RequestToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	authorizationURL, err := config.AuthorizationURL(requestToken)
	fmt.Println(authorizationURL)

	auth_redirect := GenerateRedirect(authorizationURL.String())
    auth_callback := GenerateCallback(requestSecret, config)

	http.HandleFunc("/", auth_redirect)
	http.HandleFunc("/callback", auth_callback)
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

func GenerateCallback(requestSecret string, config oauth1.Config) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        requestToken, verifier, err := oauth1.ParseAuthorizationCallback(r)
        if err != nil {
            log.Fatal("Callback: ", err)
        }
        accessToken, accessSecret, err := config.AccessToken(requestToken, requestSecret, verifier)

        fmt.Println(requestToken)
        fmt.Println(verifier)
        fmt.Println(accessToken)
        fmt.Println(accessSecret)

        token := oauth1.NewToken(accessToken, accessSecret)

        httpClient := config.Client(oauth1.NoContext, token)

        path := "https://api.twitter.com/1.1/friends/ids.json"
        resp, _ := httpClient.Get(path)

        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)

        //fmt.Println(string(body))

        var m interface{}
        json.Unmarshal(body, &m)

        fl := m.(map[string]interface{})
        ids := fl["ids"].([]interface{})

        var int_ids[]uint64

        for _, elt := range ids {
            flt := elt.(float64)
            int_ids = append(int_ids, uint64(flt))
        }

        fmt.Println(int_ids)

        http.Redirect(w, r, "http://www.google.com", http.StatusFound)
    }
}
