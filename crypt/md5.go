package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5encode md5加密
//
//@param s 待加密字符串
func (*Crypto) MD5encode(s string) string {
	instance := md5.New()
	instance.Write([]byte(s))
	sumString := instance.Sum(nil)
	return hex.EncodeToString(sumString)
}
