package main

import (
	"fmt"
	"sync"

	"github.com/qxdn/cc98sign/pkg/config"
	"github.com/qxdn/cc98sign/pkg/login"
	"github.com/qxdn/cc98sign/pkg/sign"
)

func AutoSign(user *login.User, done func()) {
	defer done()
	info := login.Login(user)
	coins := sign.SignIn(info.AccessToken)
	if coins == 0 {
		fmt.Printf("用户(%s)今日已经签到\n", user.Username)
		return
	}
	result := sign.GetSignResult(info.AccessToken)
	fmt.Printf("用户(%s)已经连续签到%d天，今日签到获得%d 98币\n", user.Username, result.LastSignInCount, coins)
}

func main() {
	configs := config.ReadConfig("config.json")
	if configs == nil {
		return
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(configs.Users))
	for _, cuser := range configs.Users {
		user := &login.User{
			Username: cuser.Username,
			Password: cuser.Password,
		}
		go AutoSign(user, waitGroup.Done)
	}
	waitGroup.Wait()
}
