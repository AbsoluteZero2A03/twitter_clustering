package util

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TokenContainer struct {
	AccessToken string `json:"access_token"`
}

func GetAccessToken(consumerKey string, consumerSecret string) string {
	consumerString := base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret))
	twitterEndpoint := "https://api.twitter.com/oauth2/token"
	client := &http.Client{}

	req, _ := http.NewRequest("POST", twitterEndpoint, bytes.NewBuffer([]byte("grant_type=client_credentials")))
	req.Header.Add("Authorization", "Basic "+consumerString)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("User-Agent", "DEEP LEARNING v0.0")
	req.Header.Add("Accept-Encoding", "gzip")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyReader, _ := gzip.NewReader(resp.Body)
	bodyText, _ := ioutil.ReadAll(bodyReader)

	var m TokenContainer
	json.Unmarshal(bodyText, &m)

	return m.AccessToken
}
