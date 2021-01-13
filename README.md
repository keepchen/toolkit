# toolkit

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/keepchen/toolkit?tab=overview)

#### 简介

封装了一些常用的工具包，逐步更新

#### 安装

```shell
go get github.com/keepchen/toolkit
```

#### 方法

* 初始化

```golang
import "github.com/keepchen/toolkit"

kit := toolkit.NewToolkit()
```

---
字符串

---

* 生成短信验证码

```golang
smsCode := kit.GenerateSMSCode(5)
```

* 生成随机字符串

```golang
shuffleStr := kit.GenerateRandomString(5)
```

* 验证手机号码

```golang
ok := kit.ValidatePhone("13890012345")
```

---
文件

---

* 解压zip文件

```golang
fileList, err := kit.Unzip("xx.zip", "/data/")
```

---
加密

---

* MD5加密

```golang
hashedStr := kit.MD5encode("123")
```

* sha256-rsa

```golang
c := kit.SetPublicKey(pubKey).SetPrivateKey(privateKey)

//签名
sign, err := c.SignUsingSha256WithRsa([]byte("test string"))

//验签
err := c.VerifySignUsingSha256WithRsa([]byte("test string"), sign)
```

---
其他

---

* gin跨域设置

```golang
...
var r *gin.Engine
r.Use(kit.StartCors("", nil))
...
```
