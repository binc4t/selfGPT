# selfGPT

[简体中文](README_zh_CN.md)

Transform into ChatGPT and surprise your friends with your impressive abilities!

<p align="center">
<img src="https://user-images.githubusercontent.com/99850403/219956652-a3815049-2058-4338-926a-b3b604c9950b.jpg"  width="540" height="1131">
</p>



## Usage 
Only run on Linux, and requires Go 1.19 or later 

Installation
```
go install github.com/anccy/selfGPT@latest
```

or build from souce
```
go build
```

Please edit the `conf.json` file and set it to your own OpenAI information. You can find your `user_name` and `access_token` at https://chat.openai.com/api/auth/session
```json
{
  "chat_gpt_user_name": "example@gmail.com",
  "chat_gpt_password": "123456",
  "chat_gpt_access_token": "xxxxxxxx"
}
```

Run
```
selfGPT
```

After executing, a two-dimensional barcode will be displayed. Scan it to log in.

## Dependence
xvfb
```
apt install xvfb
```
Chrome
```
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
apt install ./google-chrome-stable_current_amd64.deb
```

## Acknowledgment
This project is a simple version of [wechatGPT](https://github.com/yxw21/wechatgpt)

Tanks to this wechat sdk [openwechat](https://github.com/eatmoreapple/openwechat)

Have fun!
