package toolkit

import (
	"toolkit/cors"
	"toolkit/file"
	"toolkit/strings"
)

//Toolkit 结构体
type Toolkit struct {
	cors.Cors
	file.File
	strings.Strings
}

//NewToolkit 实例化
func NewToolkit() *Toolkit {
	return new(Toolkit)
}
