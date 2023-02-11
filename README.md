# selfGPT

变身 chatGPT ，让好友大吃一惊！

<p align="center">
<img src="https://user-images.githubusercontent.com/99850403/218237144-2a78afcc-76a9-46b2-9341-5feaac99e5fe.png"  width="540" height="1131">
</p>



## 使用
只能在Linux运行，需要Go 1.19或更新版本

安装
```
go install github.com/anccy/selfGPT@latest
```

或者从源码构建
```
go build
```

编辑 conf.json, 设置成你自己的openAI信息, 其中user_name 和 access_token 可以在这里找到 https://chat.openai.com/api/auth/session
```json
{
  "chat_gpt_user_name": "example@gmail.com",
  "chat_gpt_password": "123456",
  "chat_gpt_access_token": "xxxxxxxx"
}
```

运行
```
selfGPT
```

运行后会弹出二维码，扫码登陆即可



## 依赖
xvfb
```
apt install xvfb
```
Chrome
```
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
apt install ./google-chrome-stable_current_amd64.deb
```

## 致谢
这个项目是 [wechatGPT](https://github.com/yxw21/wechatgpt) 的阉割版

微信接口调用了 [openwechat](https://github.com/eatmoreapple/openwechat)

谢谢！
