# kaylyu/ministore

[wechat ministore](https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/introduction.html) development sdk written in Golang

###快速开始
```go
import github.com/kaylyu/ministore
```

### 示例
- 在获取token的时候，我这里依赖的 [github.com/silenceper/wechat](https://github.com/silenceper/wechat),当然也可以使用其他依赖方式获取
- 代码示例如下
```go
wc := wechat.NewWechat().GetOfficialAccount(&config.Config{
    AppID:     "",
    AppSecret: "",
    Cache:     cache.NewMemory(),//可使用redis替换
})

//初始化
m:=ministore.New(ministore.Config{
    AccessToken: wc.GetAccessToken(),
})

//获取运费模板
resp, err := m.GetProduct().GetFreightTemplate(nil)
fmt.Println(resp, err)

```

### License
MIT