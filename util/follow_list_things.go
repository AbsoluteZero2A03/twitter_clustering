package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type FriendsList struct {
	IdList []float64 `json:"ids"`
}

func GetFriends(id string, auth_token string) []uint64 {
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

func GetNetwork(friendsList []uint64, auth_token string) map[uint64]map[uint64]bool {
	network := make(map[uint64]map[uint64]bool)
	for _, friend := range friendsList {
		network[friend] = make(map[uint64]bool)
		idStr := strconv.FormatInt(int64(friend), 10)
		fmt.Println(idStr)
		currentFriends := GetFriends(idStr, auth_token)
		fmt.Println(currentFriends)
		for _, nextFriend := range currentFriends {
			network[friend][nextFriend] = true
		}
	}
	return network
}
