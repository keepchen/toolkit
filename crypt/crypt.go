package crypt

import "crypto/rsa"

//Crypto 结构体
type Crypto struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

//NewCrypto 实例化
func NewCrypto() *Crypto {
	return new(Crypto)
}
