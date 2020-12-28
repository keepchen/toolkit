package crypt

import "testing"

func TestMD5encode(t *testing.T) {
	str := NewCrypto().MD5encode("test string")
	if str != "6f8db599de986fab7a21625b7916589c" {
		t.Fatal("test failed")
	}
	t.Log("test passed")
}
