package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

// 小写
func Md5Encode(str string) string {
	var md = md5.New()
	md.Write([]byte(str))
	var code = hex.EncodeToString(md.Sum(nil))
	return code
}

// 大写
func MD5Encode(str string) string {
	return strings.ToUpper(Md5Encode(str))
}

// 加密

func CryptoPassword(pwd, salt string) string {
	return Md5Encode(pwd + salt)
}

// 验证密码
func ValidPassword(pwd, salt string, source_pwd string) bool {
	code := Md5Encode(pwd + salt)
	fmt.Println("code:", code)
	fmt.Println("source_pwd:", source_pwd)
	return code == source_pwd
}

var salt = string(rand.Intn(10) * 1010)

func CryptoPasswordDefaultSalt(pwd string) string {
	return Md5Encode(pwd + salt)
}
func ValidPasswordDefaultSalt(pwd, source_pwd string) bool {
	code := Md5Encode(pwd + salt)
	fmt.Println("code:", code)
	fmt.Println("source_pwd:", source_pwd)
	return code == source_pwd
}
