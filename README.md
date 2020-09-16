# toolkit

#### 简介

封装了一些常用的工具包，逐步更新

#### 安装

```
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

```
ok := kit.ValidatePhone("13890012345")
```

* MD5加密

```golang
hashedStr := kit.MD5encode("123")
```

---
文件

---

* 加压zip文件

```
fileList, err := kit.Unzip("xx.zip", "/data/")
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
