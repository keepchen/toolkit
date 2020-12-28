package toolkit

import (
	"github.com/keepchen/toolkit/cors"
	"github.com/keepchen/toolkit/crypt"
	"github.com/keepchen/toolkit/file"
	"github.com/keepchen/toolkit/strings"
)

//Toolkit 结构体
type Toolkit struct {
	cors.Cors
	file.File
	strings.Strings
	crypt.Crypto
}

//NewToolkit 实例化
func NewToolkit() *Toolkit {
	return new(Toolkit)
}
