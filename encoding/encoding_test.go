package encoding

import (
	"testing"
)

var c = &EncryptConfig{Key: "bilibili_key_GYl", IV: "biliBiliIv123456"}

const (
	plat = "13291831554"
	enc  = "dWLVOucFkBBmmXmTm6so5A=="
)

func TestEncrypt(t *testing.T) {
	s, _ := Encrypt(plat, c)
	t.Log(">>>", s)
}

func TestDecrypt(t *testing.T) {
	s, _ := Decrypt(enc, c)
	t.Log(s)

}
