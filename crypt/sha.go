package crypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

const (
	pubKeyBegin     = "-----BEGIN PUBLIC KEY-----\n"
	pubKeyEnd       = "\n-----END PUBLIC KEY-----"
	privateKeyBegin = "-----BEGIN RSA PRIVATE KEY-----\n"
	privateKeyEnd   = "\n-----END RSA PRIVATE KEY-----"
)

//SetPublicKey 设置公钥证书
func (crypt *Crypto) SetPublicKey(pubKey string) *Crypto {
	if !strings.HasPrefix(pubKey, pubKeyBegin) {
		pubKey = pubKeyBegin + pubKey
	}
	if !strings.HasPrefix(pubKey, pubKeyEnd) {
		pubKey = pubKey + pubKeyEnd
	}
	public, err := parsePublicKey([]byte(pubKey))
	if err == nil {
		crypt.publicKey = public
	}

	return crypt
}

//SetPrivateKey 设置私钥证书
func (crypt *Crypto) SetPrivateKey(privateKey string) *Crypto {
	if !strings.HasPrefix(privateKey, privateKeyBegin) {
		privateKey = privateKeyBegin + privateKey
	}
	if !strings.HasPrefix(privateKey, privateKeyEnd) {
		privateKey = privateKey + privateKeyEnd
	}
	private, err := parsePrivateKey([]byte(privateKey))
	if err == nil {
		crypt.privateKey = private
	}

	return crypt
}

//SignUsingSha256WithRsa 使用sha256-rsa算法进行签名
func (crypt *Crypto) SignUsingSha256WithRsa(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, crypt.privateKey, crypto.SHA256, d)
}

//VerifySignUsingSha256WithRsa 使用sha256-rsa算法进行签名验证
func (crypt *Crypto) VerifySignUsingSha256WithRsa(data, sign []byte) error {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)

	return rsa.VerifyPKCS1v15(crypt.publicKey, crypto.SHA256, d, sign)
}

//解析公钥
func parsePublicKey(pubKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("can not parse public key,please check it")
	}
	switch block.Type {
	case "PUBLIC KEY":
		r, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		switch r.(type) {
		case *rsa.PublicKey:
			return r.(*rsa.PublicKey), nil
		default:
			return nil, fmt.Errorf("unsupported public key type: %T", r)
		}
	default:
		return nil, fmt.Errorf("unsupported public key type: %q", block.Type)
	}
}

//解析私钥
func parsePrivateKey(privatekey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privatekey)
	if block == nil {
		return nil, errors.New("can not parse private key,please check it")
	}
	switch block.Type {
	case "RSA PRIVATE KEY":
		private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return private, nil
	default:
		return nil, fmt.Errorf("unsupported private key type: %q", block.Type)
	}
}
