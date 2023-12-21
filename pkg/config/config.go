package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	Users []User `json:"users"`
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadConfig(filepath string) *Config {
	var configs Config
	exists, _ := PathExists(filepath)
	if !exists {
		configs.Users = []User{{
			Username: "用户名",
			Password: "密码",
		}}
		file, _ := json.MarshalIndent(configs, "", " ")
		ioutil.WriteFile(filepath, file, 0644)
		fmt.Println("配置文件不存在，已经生成默认配置文件")
		return nil
	}
	jsonFile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &configs)
	return &configs
}
