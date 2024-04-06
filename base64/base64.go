package base64

import "encoding/base64"

// 把明文密码加密后返回加密过的密码
func Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// 把加密后的密码解密为明文并返回
func Decode(hashedPassword string) string {
	//密码都是用上面encode加密的，不会出错
	bytes, _ := base64.StdEncoding.DecodeString(hashedPassword)
	return string(bytes)
}
