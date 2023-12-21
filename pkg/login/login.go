package login

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 登陆用户
type User struct {
	Username string
	Password string
}

type LoginInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func Login(user *User) *LoginInfo {

	resp, err := http.PostForm("https://openid.cc98.org/connect/token", url.Values{
		"username":      {user.Username},
		"password":      {user.Password},
		"grant_type":    {"password"},
		"scope":         {"cc98-api openid offline_access"},
		"client_id":     {"9a1fd200-8687-44b1-4c20-08d50a96e5cd"}, // cc98 clientid 也可以到 https://openid.cc98.org/ 申请
		"client_secret": {"8b53f727-08e2-4509-8857-e34bf92b27f2"},
	})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		panic(err)
	}
	var info LoginInfo
	json.Unmarshal(body, &info)
	return &info
}
