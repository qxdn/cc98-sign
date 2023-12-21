package sign

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func SignIn(token string) int {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.cc98.org/me/signin", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	body := string(rawBody)
	if body == "has_signed_in_today" {
		//fmt.Println("今天已经签到过了")
		return 0
	}
	coins, _ := strconv.Atoi(body)
	return coins
}

type SignResult struct {
	LastSignInTime   string `json:"lastSignInTime"`
	LastSignInCount  int    `json:"lastSignInCount"`
	HasSignedInToday bool   `json:"hasSignedInToday"`
}

func GetSignResult(token string) *SignResult {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.cc98.org/me/signin", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var result SignResult
	json.Unmarshal(body, &result)
	return &result
}
