package main

import (
	"github.com/yxw21/chatgpt"
)

var (
	Browser *chatgpt.Browser
	Session *chatgpt.Session
	Chats   = make(map[string]*chatgpt.Chat)
)
