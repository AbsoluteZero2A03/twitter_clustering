package main

import (
	"./util"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	consumerKey := strings.TrimSpace(os.Getenv("TWITTER_APP_CONSUMER_KEY"))
	consumerSecret := strings.TrimSpace(os.Getenv("TWITTER_APP_CONSUMER_SECRET"))
	fmt.Println(consumerKey)
	fmt.Println(consumerSecret)
	fmt.Println(consumerKey + ":" + consumerSecret)

	auth_token := util.GetAccessToken(consumerKey, consumerSecret)

	req, _ := http.NewRequest("GET", "https://api.twitter.com/1.1/users/lookup.json?user_id="+os.Args[1], nil)

	req.Header.Add("Authorization", "Bearer "+auth_token)
	req.Header.Add("User-Agent", "DEEP LEARNING v0.0")

	httpClient := &http.Client{}
	resp2, _ := httpClient.Do(req)
	defer resp2.Body.Close()
	bodyText2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println(string(bodyText2))

	friends := util.GetFriends(os.Args[1], auth_token)
	fmt.Println(friends)
}
