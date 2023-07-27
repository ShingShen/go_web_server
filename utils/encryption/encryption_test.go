package encryption

import (
	"testing"
)

func TestBase64Sha256(t *testing.T) {
	password := "password123"
	expected := "75K3eLr-dx6JJFuJ7LwIpEpOFmwGZZkRiB84PURz6U8="

	result := Base64Sha256(password)
	if result != expected {
		t.Errorf("Base64Sha256(%s) = %s, expected %s", password, result, expected)
	}
}

func TestHashPassword(t *testing.T) {
	password := "password123"
	salt := "somesalt"
	expected := "6e5605d42fe720882511feecd48a6a44f2110d9d4713e1b5c4c70ed7519f9519"

	result := HashPassword(password, salt)
	if result != expected {
		t.Errorf("HashPassword(%s, %s) = %s, expected %s", password, salt, result, expected)
	}
}
