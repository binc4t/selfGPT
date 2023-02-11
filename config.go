package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	ChatGPTUserName    string `json:"chat_gpt_user_name"`
	ChatGPTPassword    string `json:"chat_gpt_password"`
	ChatGPTAccessToken string `json:"chat_gpt_access_token"`
}

func LoadConfFromFile(file string) (*Config, error) {
	conf := &Config{}
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
