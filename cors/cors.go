package cors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Cors 结构体
type Cors struct{}

//NewCors 实例化
func NewCors() *Cors {
	return new(Cors)
}

//StartCors 跨域请求处理
//
//注意：仅适用于gin-gonic/gin框架，其余框架可参考实现
//
//使用示例：
//
//var r *gin.Engine
//
//r.Use(cors.NewCors().StartCors("", nil))
func (*Cors) StartCors(origin string, headers map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var method = c.Request.Method
		if len(origin) == 0 {
			origin = c.Request.Host
			if len(origin) == 0 {
				origin = "*"
			}
		}
		if len(headers) == 0 {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT, PATCH, HEAD")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
		} else {
			for k, v := range headers {
				c.Header(k, v)
			}
		}
		//返回浏览器跨域检测
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		//处理请求
		c.Next()
	}
}
