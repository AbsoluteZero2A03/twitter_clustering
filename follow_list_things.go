package main

import (
    "fmt"
	/*"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
    "golang.org/x/oauth2"*/
    "bytes"
    "io/ioutil"
	"net/http"
	"net/http/httputil"
    "compress/gzip"
	"os"
    "strings"
    "encoding/base64"
	//"reflect"
    //"strconv"
)
/*
func getFriends(id string, config oauth1.Config, requestToken string, requestSecret string) {
    baseUrlStr := "https://api.twitter.com/1.1/friends/ids.json?user_id="
    token := oauth1.NewToken(requestToken, requestSecret)
    httpClient := config.Client(oauth1.NoContext,token)
    resp, _ := httpClient.Get(baseUrlStr + id)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Printf(string(body))
}
*/

func main() {
    consumerKey := strings.TrimSpace(os.Getenv("TWITTER_APP_CONSUMER_KEY"))
    fmt.Println(consumerKey)
    consumerSecret := strings.TrimSpace(os.Getenv("TWITTER_APP_CONSUMER_SECRET"))
    fmt.Println(consumerSecret)
    fmt.Println(consumerKey + ":" + consumerSecret)
    consumerString := base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret))

    twitterEndpoint := "https://api.twitter.com/oauth2/token"

    client := &http.Client { }
    req, _ := http.NewRequest("POST", twitterEndpoint, bytes.NewBuffer([]byte("grant_type=client_credentials")))
    req.Header.Add("Authorization", "Basic "+ consumerString)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
    req.Header.Add("User-Agent", "DEEP LEARNING v0.0")
    req.Header.Add("Accept-Encoding", "gzip")

    reqText, _ := httputil.DumpRequest(req, true)
    fmt.Println(string(reqText))

    resp, _ := client.Do(req)
    defer resp.Body.Close()
    bodyReader, _ := gzip.NewReader(resp.Body)
    bodyText, _ := ioutil.ReadAll(bodyReader)

    fmt.Println(string(bodyText))
}
