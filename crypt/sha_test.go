package crypt

import (
	"testing"
)

var (
	publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCs6cNG523lryzr031uutDCVu8u
7+SEvaymYI5wN+tj29RhqvcxTvjQ38HiXZmfjvHkW4VGhVIlgVpMqB3RQz/ZSwku
UffoAMPwEhcf+MKkcJCwCFVOC8jEMWuuyPyimYX9RWPfQ5Y4HoPTLHCLTqfLoLI9
I1Quys3SvKLwrcdKlQIDAQAB
-----END PUBLIC KEY-----`
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCs6cNG523lryzr031uutDCVu8u7+SEvaymYI5wN+tj29Rhqvcx
TvjQ38HiXZmfjvHkW4VGhVIlgVpMqB3RQz/ZSwkuUffoAMPwEhcf+MKkcJCwCFVO
C8jEMWuuyPyimYX9RWPfQ5Y4HoPTLHCLTqfLoLI9I1Quys3SvKLwrcdKlQIDAQAB
An8bsVSVBnuTTXrkSfrMwR6uAT5NRh5TJA2Cl/Q4BLyuQYbCHZ95RiyTR0LF/jec
VkhnSkEOCkwdtZnArSwZV6vh3b2U0lJs3UoPCsq2iaHdbjmgwQxlpLVL6U4tqqWH
1zUBJebDfxIFG+HXLZ4x98Pvqjq/5qOivlocb4o9i6dhAkEA1EjSMjV96LFRwUO7
H4/ZnZLLMrAKBzbfqJpgh/2JmDZ3y2fLIvMD0iGfvbGf869MNm4kc+PsrW+UcRCs
euQ6bQJBANCFXtuuJD9RiYO1s/ZDAqB6aBYa9MV5Ss/8c+RuQauIrvdzQ5hOz0bA
LM/JE2e0af8VzPGxlyOwUlo838MoN8kCQEgbWDKv0s3YdVG1ZPlKjnukohrcivt6
WyxVC9cFethWIAjaf7imXTTiMPVpQVCvya5vaThoQjuDPCPPqxavMV0CQQCohgGh
/d2hpt5Cqkllj8bBt+wDYYnNRzflXBy10z4TygXOS0OoBpmSjqIo1FNwjujTleTq
uSJzX39PGr+lVA2xAkEAzGF3Edq2tgQP/d4rQq/UhYRHPeLBQx2U5Ze0vJdo6K3P
6t3HmkwUmsm++qWpKGMG4r4VA1jmuoVNF78FaV9x1g==
-----END RSA PRIVATE KEY-----`
)

func TestSetPublicKey(t *testing.T) {
	c := NewCrypto().SetPublicKey(publicKey)
	if c.publicKey == nil {
		t.Fatal("test failed")
	}
	t.Log("test passed")
}

func TestSetPrivateKey(t *testing.T) {
	c := NewCrypto().SetPrivateKey(privateKey)
	if c.priviateKey == nil {
		t.Fatal("test failed")
	}
	t.Log("test passed")
}

func TestSignUsingSha256WithRsa(t *testing.T) {
	data := []byte("test string")
	c := NewCrypto().SetPrivateKey(privateKey).SetPublicKey(publicKey)
	_, err := c.SignUsingSha256WithRsa(data)
	if err != nil {
		t.Fatal("test failed")
	}
	t.Log("test passed")
}

func TestVerifySignUsingSha256WithRsa(t *testing.T) {
	data := []byte("test string")
	c := NewCrypto().SetPrivateKey(privateKey).SetPublicKey(publicKey)
	signByte, err := c.SignUsingSha256WithRsa(data)
	if err != nil {
		t.Fatal("test failed")
	}
	err = c.VerifySignUsingSha256WithRsa(data, signByte)
	if err != nil {
		t.Fatal("test failed")
	}
	t.Log("test passed")
}
