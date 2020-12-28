package strings

import (
	"math/rand"
	"regexp"
	"time"
)

//Strings 结构体
type Strings struct{}

//NewStrings 实例化
func NewStrings() *Strings {
	return new(Strings)
}

//GenerateSMSCode 生成短信验证码
//
//@param length 验证码长度
func (*Strings) GenerateSMSCode(length int) string {
	originString := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var code string
	if length < 1 {
		return code
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		index := rand.Intn(9) //rand(max - min) + min
		code += originString[index]
	}
	return code
}

//GenerateRandomString 生成随机字符串
//
//@param length 字符串长度
func (*Strings) GenerateRandomString(length int) string {
	originString := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z"}
	var code string
	if length < 1 {
		return code
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		index := rand.Intn(len(originString) - 1) //rand(max - min) + min
		code += originString[index]
	}
	return code
}

//ValidatePhone 验证手机号码格式
//
//@param phone 手机号
func (*Strings) ValidatePhone(phone string) bool {
	reg, err := regexp.Compile(`^1[3456789][0-9]{9}$`)
	if err != nil {
		return false
	}
	if matched := reg.Match([]byte(phone)); !matched {
		return false
	}
	return true
}
