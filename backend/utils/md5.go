package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

//Mdv
//给字符串生成md5
//@params str 需要加密的字符串
//@params salt 加密的盐
//@return str 返回md5码
func Mdv(str string, salt string) (CryptStr string) {
	h := md5.New()
	h.Write([]byte(str))
	h.Write([]byte(salt))
	st := h.Sum(nil)
	return hex.EncodeToString(st)
}

// RandomString
// returns a random string with a fixed length
// @params n return str len
// @params allowedChars
// @return str
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune
	var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
