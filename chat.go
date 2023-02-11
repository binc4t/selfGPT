package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/yxw21/chatgpt"
)

const (
	ErrMsg = "这个我无法回答，可以跟我聊点别的吗～"
)

func Run() {
	// bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册消息处理函数
	bot.MessageHandler = MessageHandler
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}

func MessageHandler(msg *openwechat.Message) {
	if msg.IsComeFromGroup() && msg.IsText() {
		sender, err := msg.Sender()
		if err != nil {
			Logger.Printf("handle msg fail, msg from: %s, content: %s, err: %v", msg.FromUserName, msg.Content, err)
			return
		}
		if sender.IsSelf() {
			return
		}
		senderID := sender.ID()
		if _, ok := Chats[senderID]; !ok {
			Chats[senderID] = chatgpt.NewChat(Browser, Session)
		}
		res, err := Chats[senderID].Send(msg.Content)
		if err != nil {
			_, _ = msg.ReplyText(ErrMsg)
			Logger.Printf("handle msg fail, msg from: %s, content: %s, err: %v", msg.FromUserName, msg.Content, err)
		} else {
			_, _ = msg.ReplyText(res.Message.Content.Parts[0])
			Logger.Printf("handle msg success, msg from: %s, content: %s, reply: %s", msg.FromUserName, msg.Content, res.Message.Content.Parts[0])
		}
	}
}
