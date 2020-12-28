package strings

import "testing"

func TestGenerateSMSCode(t *testing.T) {
	code := NewStrings().GenerateSMSCode(6)
	if len(code) != 6 {
		t.Fatal("test failed")
	}
	t.Log("test passed")
}

func TestGenerateRandomString(t *testing.T) {
	str := NewStrings().GenerateRandomString(6)
	if len(str) != 6 {
		t.Fatal("test failed")
	}
	t.Log("test passed")
}

func TestValidatePhone(t *testing.T) {
	ok := NewStrings().ValidatePhone("13890011111")
	if !ok {
		t.Fatal("test failed")
	}
	t.Log("test passed")
}
