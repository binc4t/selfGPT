package main

import (
	"github.com/yxw21/chatgpt"
	"log"
	"os"
)

var Logger *log.Logger

func initLog() {
	logFile, err := os.OpenFile("selfgpt.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	Logger = log.New(logFile, "", log.LstdFlags)
}

func main() {
	initLog()

	conf, err := LoadConfFromFile("conf.json")
	if err != nil {
		log.Fatal(err)
	}

	browser, closeBrowser, err := chatgpt.NewBrowser(chatgpt.BrowserOptions{})
	if err != nil {
		log.Fatal(err)
	}
	Browser = browser
	defer closeBrowser()

	Session = (&chatgpt.Session{
		Browser:     browser,
		Username:    conf.ChatGPTUserName,
		Password:    conf.ChatGPTPassword,
		AccessToken: conf.ChatGPTAccessToken,
	}).AutoRefresh()

	Run()
}
