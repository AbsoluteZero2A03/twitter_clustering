package util

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
    "strconv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type FriendsList struct {
	IdList []float64 `json:"ids"`
}


func getFriends(id string, auth_token string) []uint64 {
	baseUrlStr := "https://api.twitter.com/1.1/friends/ids.json?user_id="
	urlStr := baseUrlStr + id
	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header.Add("Authorization", "Bearer "+auth_token)
	req.Header.Add("User-Agent", "DEEP LEARNING v0.0")

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var fl FriendsList
	json.Unmarshal(body, &fl)
	var int_ids []uint64
	for _, elt := range fl.IdList {
		int_ids = append(int_ids, uint64(elt))
	}
	return int_ids
}

func getNetwork(friendsList []uint64, auth_token string) map[uint64]map[uint64]bool  {
    network := make(map[uint64]map[uint64]bool)
    for _, friend := range friendsList {
        network[friend] = make(map[uint64]bool)
        idStr := strconv.FormatInt(int64(friend), 10)
        currentFriends := getFriends(idStr, auth_token)
        for _, nextFriend := range currentFriends {
            network[friend][nextFriend] = true
        }
    }
    return network
}

func main() {
	consumerKey := strings.TrimSpace(os.Getenv("TWITTER_APP_CONSUMER_KEY"))
	fmt.Println(consumerKey)
	consumerSecret := strings.TrimSpace(os.Getenv("TWITTER_APP_CONSUMER_SECRET"))
	fmt.Println(consumerSecret)
	fmt.Println(consumerKey + ":" + consumerSecret)
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
	fmt.Println(m.AccessToken)

	friends := getFriends(os.Args[1], m.AccessToken)
	fmt.Println(friends)
}

